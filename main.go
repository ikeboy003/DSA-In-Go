package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	graph := NewGraph()
}
func treesCont() {

	tree := BinaryTree[int]{}

	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(8)

	tree.Insert(0)
	tree.Insert(4)
	tree.Insert(89)
	tree.Insert(95)
	tree.Insert(7)

	fmt.Println(tree.sum())

}
func WorkingWithTrees() {
	tree := newBinaryTree[int]()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(9)
	tree.Insert(8)
	tree.Insert(11)
	tree.traverseInOrder()
	fmt.Println()
	tree.traversePostOrder()
	fmt.Println()
	tree.traversePreOrder()

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
