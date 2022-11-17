package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])
}
