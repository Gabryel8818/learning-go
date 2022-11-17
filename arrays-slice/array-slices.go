package main

import "fmt"

func main() {
	// Array "Tipo e Quantidade de itens fixos"
	var array1 [5]int
	fmt.Println(array1)
	array1[0] = 12
	fmt.Println(array1[0])

	// Array com tamanho baseado na quantidade de itens passados
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	// Slice Quantidade de itens indefinidos
	pessoas := []string{}
	pessoas = append(pessoas, "Junior")
	fmt.Println(pessoas[0])

	// Arrays Internos
	slice2 := make([]float32, 10, 11)
	fmt.Println(slice2)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 6)
	// Tamanho
	fmt.Println(len(slice2))
	// Capacidade Maxima
	fmt.Println(cap(slice2))

}
