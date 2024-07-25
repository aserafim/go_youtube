package main

import (
	"fmt"
)

// Sprint pega o valor do parâmetro
// e retorna uma string com o valor
func main() {
	x := "oi"
	y := "bom dia"
	z := fmt.Sprint(x, y)

	print(z)
}
