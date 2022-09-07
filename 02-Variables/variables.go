package main

import "fmt"

func main() {
	/* Las variables se definen con "var" + el nombre  y debo pasarle el
	tipo de dato que va a almacenar. Ejm:*/
	var saludar string
	saludar = "Hola!"
	/* Tambien se puede ahorrar codigo y definirla
	 con el nombre y el valor. Ejm */
	 bridge := "ya se que oyes mis pensamientos muchacho"

	fmt.Println(saludar, "Fran", bridge)

	//Los tipos de datos:

	var b int
	var c float64
	var a string
	var d bool

	fmt.Println(a, b, c,d)


}
