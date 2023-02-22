package arrays

import "fmt"

type Slice struct {
	cap    int   `json:"cap"`
	length int   `json:"length"`
	data   []int `json:"data"`
}

func (s Slice) Size() int {
	return s.length
}

func (s Slice) Capacity() int {
	return s.cap
}

func (s Slice) IsEmpty() bool {
	return s.length == 0
}

func (s Slice) At(index int) int {
	if index > s.length || index < 0 {
		panic(fmt.Sprintf("index[%d] out of range length[%d]", index, s.length))
	}
	return s.data[index]
}

func (s Slice) resize(capacity int) {
	if capacity < s.Capacity() {
		return
	}
	s.cap = capacity
	newArray := make([]int, s.cap)
	for i, datum := range s.data {
		newArray[i] = datum
	}
	s.data = newArray
}

func (s Slice) Push(item int) {
	if s.Size() >= s.Capacity() {
		s.resize(s.Capacity() * 2)
	}
	s.data[s.Size()-1] = item
}

func (s Slice) Insert(index, item int) {
	if s.Size() >= s.Capacity() {
		s.resize(s.Capacity() * 2)
	}
	for i := s.Size() - 1; i >= index; i-- {
		s.data[i] = s.data[i+1]
	}
	s.data[index] = item
}
