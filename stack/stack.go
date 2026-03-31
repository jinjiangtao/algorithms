package stack

type stack struct {
	data []any
	len  int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(v any) {
	s.data = append(s.data, v)
	s.len++
}

func (s *stack) Pop() any {
	if s.len == 0 {
		return nil
	}
	s.len--
	return s.data[s.len]
}

func (s *stack) GetStackLen() int {
	return s.len
}
