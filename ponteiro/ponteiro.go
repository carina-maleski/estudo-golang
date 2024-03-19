package main

import "fmt"

func main() {
	var apontado int = 1
	var ponteiro *int

	ponteiro = &apontado

	fmt.Println(apontado, *ponteiro, ponteiro, &ponteiro)
	//apontado e *ponteiro imprimem o mesmo valor
	//ponteiro = endereço do apontado, &ponteiro = endereço do ponteiro :)
	
	apontado = 10
	fmt.Println(apontado, *ponteiro)
	//os dois mudam, pq tem o mesmo valor

	*ponteiro = 22
	fmt.Println(apontado, *ponteiro)
	//os dois mudam, pq tem o mesmo valor
}