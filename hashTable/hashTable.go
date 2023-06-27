package hashTable

import (
	"fmt"
	"hash/fnv"
)

/*
hash(k, m) —— m 是哈希表的大小
put(key, value) —— 如果 key 已存在则更新值
exists(key)
get(key)
remove(key)
*/

const initialCapacity = 16
const loadFactor = 0.75

type entry struct {
	key   interface{}
	value interface{}
	next  *entry
}

type HashTable struct {
	table     []*entry
	capacity  int
	threshold int
	size      int
}

func NewHashTable(capacity int) *HashTable {
	var threshold int
	if capacity == 0 {
		threshold = int(initialCapacity * loadFactor)
	} else {
		threshold = int(float64(capacity) * loadFactor)
	}
	return &HashTable{
		table:     make([]*entry, capacity),
		capacity:  capacity,
		threshold: threshold,
		size:      0,
	}
}

func (t *HashTable) hash(key interface{}, size int) int {
	// 实现全域哈希函数
	// 这里简单示例使用字符串长度乘以一个随机素数与哈希表容量取模
	// 实际应用中可以采用更复杂的全域哈希函数
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) % size
}

func (t *HashTable) resize() {
	newCapacity := t.capacity * 2
	newTable := make([]*entry, newCapacity)
	t.threshold = int(float64(newCapacity) * loadFactor)

	for _, e := range t.table {
		for e != nil {
			index := t.hash(e.key, newCapacity)
			if newTable[index] == nil {
				newTable[index] = &entry{key: e.key, value: e.value}
			} else {
				current := newTable[index]
				for current.next != nil {
					current = current.next
				}
				current.next = &entry{key: e.key, value: e.value}
			}
			e = e.next
		}
	}

	t.table = newTable
	t.capacity = newCapacity
}

func (t *HashTable) Put(key, value interface{}) {
	if t.size >= t.threshold {
		t.resize()
	}
	index := t.hash(key, t.size)
	node := t.table[index]

	if node == nil {
		t.table[index] = &entry{key: key, value: value}
		t.size++
	} else {
		for node != nil {
			if node.key == key {
				node.value = value
				return
			}
			if node.next == nil {
				break
			}
			node = node.next
		}

		node.next = &entry{key: key, value: value, next: nil}
		t.size++
	}
}

func (t *HashTable) Get(key interface{}) (value interface{}, exist bool) {
	index := t.hash(key, t.size)
	node := t.table[index]
	for node != nil {
		if node.key == key {
			return node.value, true
		}
		node = node.next
	}
	return nil, false
}

func (t *HashTable) Exists(key interface{}) bool {
	h := t.hash(key, t.size)
	node := t.table[h]
	if node == nil {
		return false
	} else {
		for node != nil {
			if node.key == key {
				return true
			}
			node = node.next
		}
		return false
	}
}

func (t *HashTable) Remove(key interface{}) {
	h := t.hash(key, t.size)
	node := t.table[h]
	var prev *entry
	for node != nil {
		if node.key == key {
			if prev != nil {
				prev.next = node.next
			} else {
				t.table[h] = node.next
			}
			return
		}
		prev = node
		node = node.next
	}
}
