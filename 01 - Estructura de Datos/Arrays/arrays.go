package main

import (
	"fmt"
)

func main() {
	//Los arrays son estaticos. Le puedo definir la longitud y tipo de dato
	var numeros [5]int

	fmt.Println("El arreglo tiene: ", numeros)

	numeros[0] = 12
	numeros[3] = 23
	fmt.Scanln(&numeros[4])
	fmt.Println("El arreglo ahora tiene: ", numeros)

	//Lo puedo definir y asignar valores:

	paises := [3]string{"Argentina", "Perú", "Chile"}
	fmt.Println(paises)

	//Puedo no definirle la longitud, pero siempre el tipo.
	continentes := [...]string{"Africa", "America"}
	//len es para saber la longitud del array
	fmt.Println(continentes, len(continentes))

	//Indices indefinidos en arreglos.
	monedas := [...]string{
		0: "Dolar",
		4: "Pesos",
	}

	fmt.Println(monedas, len(monedas))

	//Puedo tener sub-arrays extraidos del principal:

	parteDeMonedas := monedas[:2]
	fmt.Println(parteDeMonedas)

}
