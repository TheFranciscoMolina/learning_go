package main

import "fmt"

type Usuario struct {
	nombre string
	email  string
	activo bool
}

type Estudiante struct {
	user   Usuario
	codigo string
}

func main() {

	francisco := Usuario{
		nombre: "Francisco",
		email:  "fdm.molina@gmail.com",
		activo: true,
	}
	/* moli := Usuario{
		nombre: "Moli",
		email:  "fdm.moli@gmail.com",
		activo: true,
	} */

	franciscoEstudiante := Estudiante{
		user:   francisco,
		codigo: "0x23at434sdzs22",
	}

	fmt.Println(francisco)
	fmt.Println(franciscoEstudiante)

}
