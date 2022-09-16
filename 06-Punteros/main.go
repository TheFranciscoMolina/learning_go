package main

import "fmt"

func modificarValor(palabra *string) {

	fmt.Printf("%p\n", palabra)
	*palabra = "Desde la funcion!"
}

func main() {
	palabra := "Desde main!"
	fmt.Printf("%p\n", &palabra)
	fmt.Println("Antes de la fn: ", palabra)
	modificarValor(&palabra)
	fmt.Println("Despues de la fn: ", palabra)

}
