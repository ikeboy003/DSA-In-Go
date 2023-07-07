package main

import "fmt"

type Queue[T any] struct {
	queueArray []T
}

func (q *Queue[T]) enqueue(t T) {
	q.queueArray = append(q.queueArray, t)
}

func (q *Queue[T]) dequeue() (t T) {
	if len(q.queueArray) == 0 {
		return t
	}
	t = q.queueArray[0]
	q.queueArray = q.queueArray[1:]
	return t
}

func (q Queue[T]) printQueue() {

	for _, v := range q.queueArray {
		fmt.Println(v)
	}
}

func (q Queue[T]) peek() (t T) {
	if len(q.queueArray) == 0 {
		fmt.Println("Empty queue")
		return t
	}
	t = q.queueArray[0]
	return t
}
