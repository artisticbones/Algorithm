package hashTable

import (
	"fmt"
	"hash/fnv"
)

/*
type ListNode struct {
    key   interface{}
    value interface{}
    next  *ListNode
}

type HashTable struct {
    size  int
    table []*ListNode
}

func NewHashTable(size int) *HashTable {
    return &HashTable{
        size:  size,
        table: make([]*ListNode, size),
    }
}

func (ht *HashTable) hash(key interface{}) int {
    h := fnv.New32a()
    h.Write([]byte(fmt.Sprintf("%v", key)))
    return int(h.Sum32()) % ht.size
}

func (ht *HashTable) Set(key interface{}, value interface{}) {
    h := ht.hash(key)
    if ht.table[h] == nil {
        ht.table[h] = &ListNode{key, value, nil}
    } else {
        node := ht.table[h]
        for node.next != nil {
            if node.key == key {
                node.value = value
                return
            }
            node = node.next
        }
        if node.key == key {
            node.value = value
        } else {
            node.next = &ListNode{key, value, nil}
        }
    }
}

func (ht *HashTable) Get(key interface{}) interface{} {
    h := ht.hash(key)
    node := ht.table[h]
    for node != nil {
        if node.key == key {
            return node.value
        }
        node = node.next
    }
    return nil
}

func (ht *HashTable) Delete(key interface{}) {
    h := ht.hash(key)
    node := ht.table[h]
    var prev *ListNode
    for node != nil {
        if node.key == key {
            if prev != nil {
                prev.next = node.next
            } else {
                ht.table[h] = node.next
            }
            return
        }
        prev = node
        node = node.next
    }
}

*/

/*
hash(k, m) —— m 是哈希表的大小
add(key, value) —— 如果 key 已存在则更新值
exists(key)
get(key)
remove(key)
*/

type ListNode struct {
	key   interface{}
	value interface{}
	next  *ListNode
}

type HashTable struct {
	size  int
	table []*ListNode
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*ListNode, size),
	}
}

func (t *HashTable) hash(key interface{}, size int) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) % size
}

func (t *HashTable) Add(key, value interface{}) {
	h := t.hash(key, t.size)
	if t.table[h] == nil {
		t.table[h] = &ListNode{key: key, value: value, next: nil}
	} else {
		node := t.table[h]
		for node.next != nil {
			if node.key == key {
				node.value = value
				return
			}
			node = node.next
		}
		if node.key == key {
			node.value = value
		} else {
			node.next = &ListNode{key: key, value: value, next: nil}
		}
	}
}

func (t *HashTable) Get(key interface{}) (value interface{}) {
	h := t.hash(key, t.size)
	node := t.table[h]
	for node != nil {
		if node.key == key {
			value = node.value
			return
		}
		node = node.next
	}
	return nil
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
	var prev *ListNode
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
