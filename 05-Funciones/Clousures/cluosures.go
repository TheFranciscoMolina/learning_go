package main

import (
	"fmt"
	"strings"
)

// Clousure
func repetir(num int) func(palabra string) string {
	return func(palabra string) string {
		return strings.Repeat(palabra, num)
	}
}

// Recordatorio: clousures es una fn que retorna otr fn
func main() {
	veces := repetir(6)
	fmt.Println("Se repite: ", veces("Clou"))

}
