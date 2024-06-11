package main

import (
	"fmt"
)

// GoLang é uma linguagem de tipagem estática
// ou seja, criando uma var com um tipo ela
// permanece assim até o fim do código
var y = "texto"

// Ainda a atribuição usando var ou marmota :=
// faz com o que o Go já identifique o tipo
// da variável

// quando criamos uma variável
// utilizando var e seu tipo
// só podemos atribuir um valor
// em codeblock

// o valor zero é o valor
// de uma variável antes
// de inicializá-la

var z float64
var a int
var b float32
var c string
var d bool

func main() {

	z = 10.0

	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("d: %v, %T\n", d, d)
}
