package stack

type Stack struct {
	size int
	data []int
	head int
}

func NewStack(size int) *Stack {
	return &Stack{
		size: size,
		data: make([]int, 0, size),
		head: -1,
	}
}

func (s *Stack) Push(key int) {
	if s.data == nil {
		s.data = make([]int, 0, 8)
	}
	s.data = append(s.data, key)
	s.size++
	s.head++
}
func (s *Stack) Top() int {
	if s == nil {
		panic("the nil pointer, please check the data")
	}
	if s.head < 0 {
		panic("there is no item with the stack")
	}
	return s.data[s.head]
}
func (s *Stack) Pop() int {
	if s == nil {
		panic("the nil pointer, please check the data")
	}
	if s.head < 0 {
		panic("there is no item with the stack")
	}
	ret := s.data[s.head]
	s.head--
	s.size--
	return ret
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
