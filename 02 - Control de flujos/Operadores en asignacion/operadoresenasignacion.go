package main

import "fmt"

func main() {

	a := 2
	//Suma en asignacion:
	a += 2 /* (es como decir que a = a+2) */
	fmt.Println(a)
	//Resta
	a -= 1
	fmt.Println(a)
	//Multiplicacion
	a *= 4
	fmt.Println(a)
	//Dvision:
	a /= 2
	fmt.Println(a)
	//Modulo
	a %= 2
	fmt.Println(a)
}
