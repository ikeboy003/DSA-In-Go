package main

import (
	"fmt"
)

type DoublyLinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

// Extends synonmous to extends

type Comparable interface {
	int | float32 | byte
}

func (l *DoublyLinkedList[T]) AddHead(item T) {
	newItem := &Node[T]{Value: item}

	if l.Head != nil {
		l.Head.Previous = newItem
		newItem.Next = l.Head
	}

	l.Head = newItem

	if l.Tail == nil {
		l.Tail = l.Head
	}

	l.Size++

}

func (l *DoublyLinkedList[T]) AddTail(item T) {
	newItem := &Node[T]{Value: item}

	// Linking the data
	if l.Tail != nil {
		l.Tail.Next = newItem
		newItem.Previous = l.Tail
	}

	//Pointing the wrappers tail to the new Item i.e orienting
	l.Tail = newItem

	//If the list was empty l.Head is nil then the list is empty
	if l.Head == nil {
		l.Head = l.Tail
	}

	l.Size++
}

func (l DoublyLinkedList[T]) PrintForward() {

	curr := l.Head

	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}

func (l DoublyLinkedList[T]) PrintBackwards() {

	curr := l.Tail

	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Previous
	}
}

func (l DoublyLinkedList[T]) Contains(value T) bool {
	curr := l.Head

	for curr != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}

	return false
}

func (l *DoublyLinkedList[T]) Remove(item T) bool {

	curr := l.Head

	for curr != nil && curr.Value != item {
		curr = curr.Next
	}

	if curr.Next == nil {
		return false
	}

	curr.Previous.Next = curr.Next
	curr.Next.Previous = curr.Previous
	return true
}
