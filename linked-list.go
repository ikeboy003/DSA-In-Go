package main

import "fmt"

/*
LinkedList: Node with value and data, struct to hold the pointer to head
Adding the value to the pointer,
*/
type LinkedList struct {
	headNode *Node
	size     int
}

type Node struct {
	data         int
	previousNode *Node
	nextNode     *Node
	index        int
}

func (l *LinkedList) appendToTail(value int) {
	endNode := &Node{
		data:     value,
		nextNode: nil,
	}
	if l.headNode == nil {
		l.headNode = endNode
		l.headNode.index = l.size
		l.size++
	} else {

		// keep looping til nil then set the last node = end node
		curr := l.headNode
		for curr.nextNode != nil {
			curr = curr.nextNode
		}
		curr.nextNode = endNode
		curr.nextNode.index = l.size
		l.size++
	}

}

func (list *LinkedList) Print() {

	current := list.headNode
	for current != nil {
		fmt.Println(current.index, current.data)
		current = current.nextNode
	}
}

/*
Can also be implement with just the node, but problematic if head changes
*/

func (n *Node) appendToTailWithJustNode(value int) {
	endNode := &Node{
		data:     value,
		nextNode: nil,
	}

	for n.nextNode != nil {
		n = n.nextNode
	}
	n.nextNode = endNode
}

/*
o(n) complexity for looping through a singly linked list
*/
func (l *LinkedList) contains(val int) bool {
	curr := l.headNode
	for curr != nil {
		if curr.equals(val) {
			return true
		}
		curr = curr.nextNode
	}
	return false
}

func (n *Node) printNode() {
	curr := n
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.nextNode
	}
}

func (n Node) equals(val int) bool {

	return n.data == val
}
