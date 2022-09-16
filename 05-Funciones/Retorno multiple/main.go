package main

import "fmt"

//Usando ... decimos que los argumentos son indefinidos
//Si pasamos 2 argumentos, primero van los determinados
func sumar(nombre string, numeros ...int) (string, int) {
	var total int
	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	//Ignoro el indice con _,
	//y uso el valor num
	for _, num := range numeros {
		total += num
	}
	return mensaje, total
}

func main() {
	mensaje, total := sumar("Francisco", 5+5+5)
	fmt.Println(mensaje, total)
}
