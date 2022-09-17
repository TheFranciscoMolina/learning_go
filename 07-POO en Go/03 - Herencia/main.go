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

//Herencia:
type Invitado struct {
	Usuario
	promedio float64
}

func main() {

	//Creo objetos de esa estructura:
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

	//Objetos y herencia:

	invitadoUno := Invitado{
		promedio: 125.3,
	}

	invitadoUno.nombre = "User1"
	invitadoUno.contraseña = "qwerty"

	fmt.Println(invitadoUno)

}
