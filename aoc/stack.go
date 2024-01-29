package aoc

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}
