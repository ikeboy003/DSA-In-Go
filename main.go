package main

import (
	"math/rand"
	"time"
)

func main() {
	tree := newBinaryTree[int]()
}

func workingWithQueues() {

	que := Queue[int]{}

	que.Push(1)
	que.Push(2)
	que.Push(3)
	que.Push(4)
	que.Push(5)
	que.Pop()
	que.Print()
}
func TestingSortedLinkedList() {
	SLL := LinkedList[int]{}
	DLL := DoublyLinkedList[int]{}
	n := 10
	for n > 0 {
		SLL.AddHead(n)
		DLL.AddHead(n)
		n--
	}

	Sorted := SortedList[int]{}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 14; i++ {
		Sorted.AddSorted(rand.Intn(100))
	}

	Sorted.PrintForward()

}
