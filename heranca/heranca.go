package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	pessoa1 := pessoa{"Joao", "Pedro", 15, 1.78}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "TI"}
	fmt.Println(estudante1.nome)

}
