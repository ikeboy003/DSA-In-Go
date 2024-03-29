package main

import (
	"fmt"
)

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Peek() (t T) {
	if q.isEmpty() {
		return
	}
	return q.arr[0]
}

func (q *Queue[T]) isEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue[T]) Push(item T) {

	q.arr = append(q.arr, item)
}

func (q *Queue[T]) Pop() (t T) {
	if q.isEmpty() {
		return
	}

	t = q.arr[0]

	q.arr = q.arr[1:]
	return t

}

func (q *Queue[T]) Print() {
	for _, v := range q.arr {
		fmt.Println(v)
	}

}

func (q *Queue[T]) length() int {
	return len(q.arr)
}
