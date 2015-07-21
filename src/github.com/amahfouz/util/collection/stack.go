package collection

// trivial implementation of a stack of int

type Stack struct {
	stack []int
	size  int
}

func NewStack(initCapacity int) Stack {
	return Stack{make([]int, 0, initCapacity), 0}
}

func (s Stack) Size() int {
	return s.size
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(element int) {
	if s.size == len(s.stack) {
		s.stack = append(s.stack, element)
	} else {
		s.stack[s.size] = element
	}

	s.size++
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Pop called on an empty stack.")
	} else {
		element := s.stack[s.size-1]
		s.size--
		return element
	}
}

func (s *Stack) Clear() {
	s.size = 0
}
