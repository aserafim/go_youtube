package main
import ("fmt")

type tipo int
//var x tipo
var x tipo = 42
func main(){
	fmt.Printf("%T %v \n", x,x)

}