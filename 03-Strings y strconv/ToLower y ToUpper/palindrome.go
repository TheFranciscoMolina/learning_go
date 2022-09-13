package main

import (
	"fmt"
	"strings"
)

func rev(texto string) string {
	arrTexto := strings.Split(texto, "")
	arrRev := make([]string, 0)

	for i := len(arrTexto) - 1; i >= 0; i-- {
		arrRev = append(arrRev, arrTexto[i])
	}

	return strings.Join(arrRev, "")
}

func esPalindromo(palabra string) bool {
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")
	palabraRev := rev(palabra)
	fmt.Println(palabraRev)
	return palabra == palabraRev
}

func main() {
	if esPalindromo("oso") {
		fmt.Println("Es palindromo!")
	} else {
		fmt.Println("No es palindromo")
	}
}
