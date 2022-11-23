package main

import (
	"fmt"
	//	"time"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }
	//
	// for j := 0; j <= 10; j++ {
	// 	fmt.Println(j)
	// }

	user := map[string]interface{}{
		"Nome":      "Gabreu",
		"Sobrenome": "Melo",
	}

	for nome, sobrenome := range user {
		fmt.Printf("Chave: %s\n Valor: %s \n", nome, sobrenome)
	}
}
