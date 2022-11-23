package main

import (
	"fmt"
	"math"
)

// criacao do retangulo como struct
type retangulo struct {
	altura  float64
	largura float64
}

// implementacao do metodo ao struct retangulo que se chama area e retorna um float64
func (ret retangulo) area() float64 {
	return ret.altura * ret.largura
}

// criacao da struct circulo
type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * (math.Pow(c.raio, 2))
}

// criacao da interface
// a mesma possibilita a flexibilidade de tipos
// nesse caso ao usar essa interface seu struct obrigatoriamente precisa
// conter um metodo chamado area que retorna um float64
type forma interface {
	area() float64
}

// funcao que receber a inteface e retorna o valor de area da mesma
func escreverArea(f forma) {
	fmt.Printf("area da forma: %0.2f", f.area())
}

func main() {
	// passando valores para o retangulo
	r := retangulo{10, 15}
	//retornando o resultado de 10*15
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
