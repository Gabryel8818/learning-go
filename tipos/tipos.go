package main

import "fmt"

func main() {
	var num1 int8 = 1
	var num2 int16 = 2
	var num3 int32 = 3
	var num4 int64 = 4

	fmt.Println(num1, num2, num3, num4)

	var num5 uint = 10 // Unsigned int
	fmt.Println(num5)

	// alias
	// INT32 = rune
	var num6 rune = 1231
	fmt.Println(num6)

	// BYTE = uint8
	var num7 byte = 123
	fmt.Println(num7)
}
