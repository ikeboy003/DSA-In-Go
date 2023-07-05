package main

import (
	"fmt"
)

type DoublyLinkedList struct {
	size     int
	headNode *Node
	tailNode *Node
}

// Adding to head o(1)
func (dl *DoublyLinkedList) addToHead(value int) {
	var p2newNode *Node = &Node{data: value, nextNode: dl.headNode}
	if dl.headNode == nil {
		dl.headNode = p2newNode
		dl.tailNode = p2newNode
		dl.size++
	} else {

		dl.headNode.previousNode = p2newNode
		p2newNode.nextNode = dl.headNode
		dl.headNode = p2newNode

	}

}
func (dl *DoublyLinkedList) addToTail(value int) {
	var p2newNode *Node = &Node{data: value}
	if dl.tailNode == nil {
		dl.headNode = p2newNode
		dl.tailNode = p2newNode
		dl.size++
	} else {

		dl.tailNode.nextNode = p2newNode
		p2newNode.previousNode = dl.tailNode
		dl.tailNode = p2newNode

	}

}
func (dl *DoublyLinkedList) print() {

	curr := dl.headNode

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.nextNode
	}
}
func getDecimalValue(head *Node) int {
	result := 0

	current := head
	for current != nil {

		result = (result << 1) + current.data
		fmt.Println(result, current.data)
		current = current.nextNode
	}

	return result
}
