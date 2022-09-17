package main

import "fmt"

type Usuario struct {
	nombre     string
	contraseña string
}

func main() {
	nuevoUsuario := Usuario{"Fran", "asd123"}
	fmt.Println(nuevoUsuario)

	otroUsuario := Usuario{
		nombre:     "Moli",
		contraseña: "asd123asd",
	}

	fmt.Println(otroUsuario)
}
