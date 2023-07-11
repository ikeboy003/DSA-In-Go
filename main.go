package main

func main() {
	SLL := LinkedList[int]{}
	DLL := DoublyLinkedList[int]{}
	n := 10
	for n > 0 {
		SLL.AddHead(n)
		DLL.AddHead(n)
		n--
	}

	SLL.Remove(10)
	SLL.PrintForward()

}
