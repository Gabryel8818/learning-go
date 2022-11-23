package main

import "fmt"

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("RECUPERANDOOO MLK")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recovery()
	media := (n1 + n2) / 2

	if media <= 5 {
		panic("MLK TU REPETIU")
	} else {
		return true
	}

}

func main() {

	alunoEstaAprovado(1, 1)
}
