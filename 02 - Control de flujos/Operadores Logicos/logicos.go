package main

import (
	"fmt"
)

func main() {
	//En go tenemos tres operadores logicos:

	//Not (!): va a negar:

	fmt.Println(!true)

	//And (&&): Deben darse las dos o mas condiciones:
	fmt.Println(true && false)
	//Or(||): tiene que darse una u otra condicion:
	fmt.Println(true || false)
}
