package main

import "fmt"

//EScructura o "Strict"
type Usuario struct {
	nombre     string
	contraseña string
}

//Metodos:

func (change *Usuario) changePassword(newPassword string) {
	change.contraseña = newPassword
	fmt.Println("La nueva contraseña es:", change.contraseña)
}

func main() {
	nuevoUsuario := Usuario{"Fran", "asd123"}
	fmt.Println(nuevoUsuario)

	otroUsuario := Usuario{
		nombre:     "Moli",
		contraseña: "asd123asd",
	}
	fmt.Println(otroUsuario)
	//Uso el metodo
	otroUsuario.changePassword("nuevopass")

	fmt.Println(otroUsuario)
}
