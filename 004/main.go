package main

import (
	"fmt"
)

// Criamos um tipo hotdog
// que tem como tipo base o int
// o tipo subjacente 
type hotdog int
var b hotdog

func main() {
	fmt.Printf("b: %T", b)
}
