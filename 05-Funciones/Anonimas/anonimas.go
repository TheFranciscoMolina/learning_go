package main

import "fmt"

func main() {

	/* 	func() {
		fmt.Println("Soy una fn anonima")
	}() */

	funAnonima := func() {
		fmt.Println("Soy una fn anonima")
	}

	funAnonima()

	funAnonimaConNombre := func(nombre string) string {
		return fmt.Sprintf("Soy una fn anonima con el nombre %s", nombre)
	}

	fmt.Println(funAnonimaConNombre("Fran"))
}
