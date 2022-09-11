package main

import "fmt"

func main() {

	a := 4
	b := 6

	var r bool

	//Iguales:
	r = a == b
	fmt.Printf("El valor de %d es igual al de %d? --> %t\n", a, b, r)

	//Distintos
	r = a != b
	fmt.Printf("El valor de %d es distinto al de %d? --> %t\n", a, b, r)
	//Mayor que
	r = a > b
	fmt.Printf("El valor de %d es mayor al de %d? --> %t\n", a, b, r)
	//Menor que
	r = a < b
	fmt.Printf("El valor de %d es menor al de %d? --> %t\n", a, b, r)
	//Mayor o igual que
	r = a >= b
	fmt.Printf("El valor de %d es mayor o igual al de %d? --> %t\n", a, b, r)
	//Menor o igual que
	r = a <= b
	fmt.Printf("El valor de %d es menor o igual al de %d? --> %t\n", a, b, r)
}
