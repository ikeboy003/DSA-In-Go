package main

import (
	"fmt"
)

type LinkedList[T comparable] struct {
	Head *Node[T]
}

func (l *LinkedList[T]) AddHead(item T) {
	newItem := &Node[T]{Value: item}

	if l.Head != nil {
		newItem.Next = l.Head
	}

	l.Head = newItem

}

func (l *LinkedList[T]) PrintForward() {

	curr := l.Head

	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}

func (l *LinkedList[T]) AddTail(item T) {

	newItem := &Node[T]{Value: item}

	if l.Head != nil {
		curr := l.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newItem
	} else {
		l.Head = newItem
	}

}

func (l *LinkedList[T]) Remove(item T) {

	curr := l.Head
	prev := &Node[T]{}

	for curr != nil && curr.Value != item {
		prev = curr
		curr = curr.Next
	}

	if curr == nil {
		fmt.Println("Not FounD")
	} else {

		prev.Next = curr.Next
	}

}

func (l LinkedList[T]) Contains(value T) bool {
	curr := l.Head

	for curr != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}

	return false
}
