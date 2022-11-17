package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func main() {
	fmt.Println(somar(100, 2))

	f := func(txt string) string {
		return txt
	}

	result := f("ADADAD")
	fmt.Println(result)
}
