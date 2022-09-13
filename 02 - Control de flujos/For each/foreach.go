package main

import "fmt"

func main() {
	paises := [...]string{"Argentina", "Colombia", "Mexico"}

	for index, data := range paises {
		fmt.Println(index, data)
	}

	//usando una "_" niego recorrer la clave o el valor
	for _, data := range paises {
		fmt.Println(data)
	}
}
