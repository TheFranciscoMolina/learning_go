package main

import (
	"fmt"
	"paquetes/models"
	// "paquetes/saludar"
)

func main() {
	// saludar.Hola()
	persona1 := models.Persona{}
	persona1.Constructor("Francisco", 30)
	fmt.Println(persona1)

}
