package main

import "fmt"

func main() {
	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel2, variavel)

	variavel++
	fmt.Println(variavel2, variavel)

	// Ponteiro
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(*ponteiro)

}
