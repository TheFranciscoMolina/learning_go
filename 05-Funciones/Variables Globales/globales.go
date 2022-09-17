package main

import "fmt"

//Variable global con var
var saludo string

func hola() {
	saludo := "Hola!"
	fmt.Println(saludo)
}

func adios() {
	saludo := "Adios!"
	fmt.Println(saludo)
}

func main() {
	saludo := "Saludos desde main!"
	//Usando "defer" hago que esa funcion se ejecute al final
	defer hola()
	adios()
	fmt.Println(saludo)
}
