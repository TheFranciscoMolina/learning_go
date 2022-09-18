package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {
		num, _ := strconv.Atoi(valor)
		result += num
	}
	return result
}

func main() {
	var expresion string
	var result int

	fmt.Println("***Ingresa el calculo***")
	fmt.Scanln(&expresion)

	result = sumar(expresion)
	println("El resultado es: ", result)

}
