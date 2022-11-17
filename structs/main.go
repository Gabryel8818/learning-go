package main

import "fmt"

func main() {
	type endereco struct {
		rua    string
		numero int
	}
	type usuario struct {
		nome     string
		idade    int
		endereco endereco
	}

	endereco1 := endereco{"RUa nic", 655}
	usuario1 := usuario{"Junior", 32, endereco1}
	fmt.Printf("Nome: %s \n Idade: %d \n Rua: %s", usuario1.nome, usuario1.idade, endereco1.rua)
}
