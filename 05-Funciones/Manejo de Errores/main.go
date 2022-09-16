package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre cero. Ingresa otro numero")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {
	resultado, error := dividir(10, 2)

	if error == nil {
		fmt.Println("El resultado es:", resultado)
	} else {
		fmt.Println(error)
	}
}
