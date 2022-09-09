package main

import "fmt"

func main() {

	//Son similares a los objetos literales de Javascript. Tienen clave y valor
	//Declaro la var, uso la fn make
	//Creo un "mapa" y se asigna primero el tipo de valor que lleva la clave, y despues el tipo del valor
	dias := make(map[int]string)

	fmt.Println(dias)

}
