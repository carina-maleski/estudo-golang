package main

import "fmt"

func main() {
	fmt.Println(somaInt(1, 2))
	fmt.Println(somaFloat(10, 2))
}

//criando funções = método em java
func somaInt(a int, b int) int {
	return a + b
}

func somaFloat(a float64, b float64) float64 {
	return a + b
}

//alguns tipos de dados em go
// > string
// > int int32 int64 
// > uint uint32 uint64 (unsigned => apenas inteiros positivos)
// > float32 float64
// > bool
// > byte
// > rune
// > any | interface