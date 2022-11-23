package main

import "fmt"

func soma(numeros ...int) int {

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {

	fmt.Println(soma(200, 100, 10, 20, 30))

}
