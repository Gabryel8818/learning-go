package main

import "fmt"

func Calculos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2

	return
}

func main() {

	fmt.Println(Calculos(100, 50))
}
