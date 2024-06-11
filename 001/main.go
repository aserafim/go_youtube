package main

import (
	"fmt"
)

var z = "Olá, bom dia." //usado apenas para variáveis de package leve scope

func main() {
	_, erros := fmt.Println("Hello, World", 100, "com números")
	fmt.Println(erros)

	x := 10
	y := "Isso é uma string"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	fmt.Printf("z: %v\n", z)

}
