package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("oi")
	generic(1)
	generic(false)

	//println recebe interfaces, e de qualquer tipo
	// esse seria um exemplo de interface generica
}
