package main

import "fmt"

func closure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	text := "Dentro da func main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}
