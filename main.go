package main

import "fmt"

func main() {

	fmt.Println(1 << 2)

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
