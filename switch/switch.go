package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda Feira"
	default:
		return "Numero invalido"
	}
}

func main() {

	fmt.Println(diaDaSemana(2))

}
