package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) push(item T) {

	s.items = append(s.items, item)

}
func (s *Stack[T]) pop() (t T) {
	if len(s.items) == 0 {
		return t
	}
	index := len(s.items) - 1
	t = s.items[index]
	s.items = s.items[:index]
	return t
}

func (s Stack[T]) printStack() {
	fmt.Println("Printing STack")
	for _, v := range s.items {
		fmt.Println(v)
	}
}
