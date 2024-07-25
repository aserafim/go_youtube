package main
import ("fmt")

type tipo int
var y tipo
func main(){
	fmt.Printf("Tipo: %T\nValor: %v\n", y, y)
	y = 42
	fmt.Println(y)
	//Realizando a conversão

	x := int(y)
	fmt.Printf("Tipo de x: %T\nValor de x: %v\n", x, x)
}