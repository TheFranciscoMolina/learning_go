package main

import "fmt"

func main() {
	nombre := "Francisco"
	var apellido string
	edad := 30

	fmt.Printf("Hola! Me lllamo %s y tengo %d años. \n", nombre, edad)
	//\n es para hacer salto de lineas
	//%s es el tipo de dato string y el %d es int
	//Si no se sabe el tipo de dato que se espera podemos usar "%v"
	fmt.Printf("Hola! Me lllamo %s y tengo %d años. \n", nombre, edad)

	//Tambien con Pintf se puede saber el tipo de dato de una variable usando %T:

	fmt.Printf("Edad es del tipo: %T", edad)

	//Sprintf formatea la informacion para retornarla
	mensaje := fmt.Sprintf("Hola! Me lllamo %s y tengo %d años", nombre, edad)
	fmt.Println("La variable mensaje dice: ", mensaje)

	//Para pedir input por la consola:

	fmt.Print("Ingresa tu apellido:")
	fmt.Scanln(&apellido)

	fmt.Println("El apellido es: ", apellido)

}
