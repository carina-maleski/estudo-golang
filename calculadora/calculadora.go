package main

import "fmt"

type Calculadora struct {
	Numero1 int
	Numero2 int
}

func (c Calculadora) Soma() int {
	return c.Numero1 + c.Numero2
}

func (c Calculadora) Sub() int {
	return c.Numero1 - c.Numero2
}

func (c Calculadora) Div() int {
	return c.Numero1 / c.Numero2
}

func (c Calculadora) Mult() int {
	return c.Numero1 * c.Numero2
}

func (c Calculadora) Resto() int {
	return c.Numero1 % c.Numero2
}

func main() {
	c := Calculadora{Numero1: 200, Numero2:4}
	fmt.Println(c)
	fmt.Println(c.Soma(), c.Sub(), c.Mult(), c.Div(), c.Resto())
}