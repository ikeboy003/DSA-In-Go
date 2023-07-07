package main

import "fmt"

func main() {

	queue := Queue[int]{}
	queue.enqueue(1)
	queue.enqueue(3)
	queue.dequeue()
	queue.printQueue()
	fmt.Println(queue.peek())
	queue.dequeue()
	fmt.Println(queue.peek())

	myMap := HashMap[int, string]{
		data: make(map[int]string),
	}

	myMap.Put(1, "2")
	val, _ := myMap.Get(1)
	fmt.Println(val)

}

func bitShiftingExample(bin int) int {

	var dec, toPower int

	for bin > 0 {
		num := bin % 10
		out := num << toPower
		fmt.Println(out)
		dec += out
		toPower++
		bin /= 10
	}

	return dec

}
