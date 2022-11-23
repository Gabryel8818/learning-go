package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalComPont(numero *int) {
	*numero = *numero * -1

}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNum := 30
	inverterSinalComPont(&novoNum)
	fmt.Println(novoNum)
	fmt.Println(novoNum)
}
