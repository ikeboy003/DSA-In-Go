package main

import "fmt"

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Peek() T {
	return s.arr[0]
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack[T]) Push(item T) {

	s.arr = append(s.arr, item)
}

func (s *Stack[T]) Pop() (t T) {
	if s.isEmpty() {
		return t
	}
	length := len(s.arr) - 1
	t = s.arr[length]

	s.arr = s.arr[:length]
	return t

}

func (s *Stack[T]) Print() {
	for _, v := range s.arr {
		fmt.Println(v)
	}

}
