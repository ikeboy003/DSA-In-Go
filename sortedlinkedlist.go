package main

import "fmt"

type SortedList[T Comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func (s *SortedList[T]) AddSorted(item T) {
	newItem := &Node[T]{Value: item}

	if s.Head == nil && s.Tail == nil {
		s.Head, s.Tail = newItem, newItem

	} else {
		if item <= s.Head.Value {
			s.Head.Previous = newItem
			newItem.Next = s.Head
			s.Head = newItem
		} else if item >= s.Tail.Value {
			s.Tail.Next = newItem
			newItem.Previous = s.Tail
			s.Tail = newItem
		} else {
			curr := s.Head
			for curr != nil && curr.Value < item {
				curr = curr.Next
			}

			newItem.Next = curr
			newItem.Previous = curr.Previous

			curr.Previous.Next = newItem
			curr.Previous = newItem

		}
	}

}

func (s SortedList[T]) PrintForward() {

	curr := s.Head

	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}
