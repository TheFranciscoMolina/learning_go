package main

import "fmt"

func main() {

	//Son similares a los objetos literales de Javascript. Tienen clave y valor
	//Declaro la var, uso la fn make
	//Creo un "mapa" y se asigna primero el tipo de valor que lleva la clave, y despues el tipo del valor
	dias := make(map[int]string)

	fmt.Println(dias)

	//Agrego los datos a ese mapa:
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[6] = "Sabado"
	dias[3] = "Miercoles"
	dias[5] = "Viernes"
	dias[4] = "Jueves"
	fmt.Println(dias)

	//Quito el valor del mapa usando delete:

	delete(dias, 2)
	fmt.Println(dias)
	fmt.Println("La longitud del mapa es: ", len(dias))

	//Mapa un póco mas complejo:
	//Por ejm, una clave que tiene como valor un arreglo:

	estudiantes := make(map[string][]int)
	fmt.Println(estudiantes)

	//Le pasamos valores:
	estudiantes["Francisco"] = []int{22, 25, 30}
	estudiantes["Daniel"] = []int{4, 21, 5}
	fmt.Println(estudiantes)

	//Imprimo solo un array:
	fmt.Println("Los datos de Fran son: ", estudiantes["Francisco"])

	//imprimo solo un elemento de ese array:
	fmt.Println("Solo un elemento de la lista: ", estudiantes["Francisco"][0])
}
