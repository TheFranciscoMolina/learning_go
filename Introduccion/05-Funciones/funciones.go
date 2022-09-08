package main

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func main() {
	result := sumar(12, 25)
	fmt.Println("El resultado es:", result)
}
