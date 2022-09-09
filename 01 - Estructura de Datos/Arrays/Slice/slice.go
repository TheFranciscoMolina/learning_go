package main

import "fmt"

func main() {
	//Slicen
	numeros := []int{1, 2, 3}
	fmt.Println("Sin slice: ", numeros)
	numeros = append(numeros, 6, 5, 4)
	fmt.Println("Con slice: ", numeros)
	//Sub slice, modifico y guardo el valor pero no afecta a la variable padres
	subNumeros := numeros[0:2]
	//Por ejm, vuelvo a modificar el indice 0 del primer slice
	numeros[0] = 22
	fmt.Println("Con el sub-slice es", subNumeros)

	//Punteros (ejm: [0:2])
	//Longitud (ejm: len(numeros))
	//Capacidad (ejm: 6)

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
	meses = append(meses, "Abril", "Mayo", "Junio")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
