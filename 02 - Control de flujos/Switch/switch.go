package main

import "fmt"

func main() {
	var vocal string
	fmt.Print("Ingresa una letra:")
	fmt.Scanln(&vocal)

	/* switch {
	case vocal == "a" || vocal == "A":
		fmt.Println(vocal, "es vocal")
	case vocal == "e" || vocal == "E":
		fmt.Println(vocal, "es vocal")
	case vocal == "i" || vocal == "I":
		fmt.Println(vocal, "es vocal")
	case vocal == "o" || vocal == "O":
		fmt.Println(vocal, "es vocal")
	case vocal == "u" || vocal == "U":
		fmt.Println(vocal, "es vocal")
	default:
		fmt.Println("Le letra", vocal, "no es una vocal")
	} */

	//De esta manera busca a "vocal" en los casos y ahorro codigo
	switch vocal {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println(vocal, "es vocal")
	default:
		fmt.Println("Le letra", vocal, "no es una vocal")
	}

}
