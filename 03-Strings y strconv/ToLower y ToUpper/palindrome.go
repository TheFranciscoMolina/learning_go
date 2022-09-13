package main

import (
	"fmt"
	"strings"
)

func esPalindromo(palabra string) {
	palabra = strings.ToLower(palabra)

	fmt.Println(palabra)
}

func main() {
	esPalindromo("ya sabia")
}
