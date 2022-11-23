package main

import "fmt"

type user struct {
	nome  string
	idade int
}

// metodo em go
func (u user) salvar() { // primeiro campo eh o atributo e o segundo o nome do metodo
	fmt.Printf("Salvando os dados do user: %s \n", u.nome)

}

func (u *user) aniversario() {
	u.idade++
	fmt.Println(u.idade)
}

// metodo retorna true de acordo com a idade
func (u user) maioridade() bool {
	return u.idade >= 18
}

func main() {

	usuario := user{"junim", 10}
	fmt.Println(usuario)
	usuario.salvar()

	fmt.Println(usuario.maioridade())
	usuario.aniversario()
}
