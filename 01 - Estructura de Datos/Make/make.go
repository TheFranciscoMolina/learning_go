package main

import "fmt"

func main() {

	numeros := make([]int, 2, 4)

	numeros[0] = 100
	numeros[1] = 200
	numeros = append(numeros, 500)

	fmt.Println(numeros)
}
