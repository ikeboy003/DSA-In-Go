package main

import "fmt"

func main() {
	bitShiftingExample(101101)

}

func bitShiftingExample(bin int) int {

	temp := bin
	var num, dec, toPower int
	for bin > 0 {
		num = bin % 10
		out := num << toPower
		fmt.Println(num, out)
		dec = dec + num
		toPower++
		bin /= 10
	}
	return temp
}
