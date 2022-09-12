package main

import (
	"fmt"
)

func main() {

	/* Si la condicion se cuenple
	pasa al if directamente */
	if nombre, edad := "Moli", 30; nombre == "Fran" {
		fmt.Println("Bienvenido de nuevo, ", nombre)
	} else {
		fmt.Printf("Nombre: %s, Edad: %d \n", nombre, edad)
	}

	//Obtener los valores de un mapa:
	user := make(map[string]string)

	user["Francisco"] = "fdm.molina@gmail.com"
	user["Daniel"] = "fdm_molina@hotmail.com"

	fmt.Println(user)

	//Verificar si el valor de un mapa existe y evitar ver solo el render del espacio vacio:
	correo, ok := user["Francisco"]
	fmt.Println(correo, ok)

	//A esto lo puedo usar como flag con una variable if:

	if mail, ok := user["Francisco"]; ok {
		fmt.Printf("El correo de Francisco es: %s \n", mail)
	} else {
		fmt.Println("No se pudo acceder al correo")
	}

}
