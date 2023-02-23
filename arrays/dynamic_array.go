package arrays

import "fmt"

type Slice struct {
	cap    int
	length int
	data   []int
}

func NewSlice(length, capacity int) *Slice {
	if length > capacity {
		panic("length cannot bigger than capacity!")
	}
	data := make([]int, capacity)
	for i := 0; i < length; i++ {
		data[i] = 0
	}
	return &Slice{
		cap:    capacity,
		length: length,
		data:   data,
	}
}

func (s *Slice) Size() int {
	return s.length
}

func (s *Slice) Capacity() int {
	return s.cap
}

func (s *Slice) IsEmpty() bool {
	return s.length == 0
}

func (s *Slice) At(index int) int {
	if index > s.length || index < 0 {
		panic(fmt.Sprintf("index[%d] out of range length[%d]", index, s.length))
	}
	return s.data[index]
}

func (s *Slice) resize(capacity int) {
	if capacity < 0 {
		panic("resize error")
	}
	change := false
	if capacity >= s.Capacity() {
		capacity = s.Capacity() * 2
		change = true
	}
	if capacity <= s.Capacity()/4 {
		capacity = s.Capacity() / 2
		change = true
	}
	if !change {
		return
	}
	s.cap = capacity
	newArray := make([]int, s.cap)
	copy(newArray, s.data)
	s.data = newArray
}

func (s *Slice) Push(item int) {
	if s.Size() >= s.Capacity() {
		s.resize(s.Size())
	}
	s.data[s.Size()] = item
	s.length++
}

func (s *Slice) Insert(index, item int) {
	if s.Size() >= s.Capacity() {
		s.resize(s.Size())
	}
	for i := s.Size() - 1; i >= index; i-- {
		s.data[i+1] = s.data[i]
	}
	s.data[index] = item
	s.length++
}

func (s *Slice) Pop() int {
	ret := s.data[s.Size()-1]
	s.data[s.Size()-1] = s.data[s.Size()]
	s.length--
	defer s.resize(s.Size())
	return ret
}

func (s *Slice) Delete(index int) {
	copy(s.data[index:], s.data[index+1:])
}

func (s *Slice) Remove(item int) (idxs []int) {
	for i := s.Size() - 1; i >= 0; i-- {
		if s.data[i] != item {
			continue
		} else {
			if s.Size() == s.Capacity() {
				s.data = s.data[:i]
			} else {
				copy(s.data[i:], s.data[i+1:])
			}
			idxs = append(idxs, i)
			s.length--
		}
	}
	defer s.resize(s.Size())
	return idxs
}

func (s *Slice) Find(item int) (idx int) {
	idx = -1
	for i, e := range s.data {
		if e == item {
			idx = i
			break
		}
	}
	return idx
}

func (s *Slice) GetArray() []int {
	return s.data[:s.Size()]
}
