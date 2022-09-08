package main

import "fmt"

func main() {
	//Slicen
	numeros := []int{2, 3, 5}
	fmt.Println("Sin slice: ", numeros)
	numeros = append(numeros, 55, 22, 32)
	fmt.Println("Con slice: ", numeros)
}
