package main

import "fmt"

func main() {

	//En GO no existe el "while" pero podemos imitarlo asi.

	numeros := 26532
	flag := 0

	for numeros > 0 {
		numeros /= 10
		flag++
	}

	fmt.Println("En total hay", flag, "digitos.")

	//Tipico bucle "for":

	for i := 0; i <= 10; i++ {
		fmt.Println("Esta es la vuelta n°", i)
	}

	for i := 50; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("El n°", i, "es par.")
		}
	}
}
