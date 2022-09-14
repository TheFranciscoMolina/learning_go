package main

import "fmt"

//Usando ... decimos que los argumentos son indefinidos
func sumar(numeros ...int) int {
	var total int

	//Ignoro el indice con _,
	//y uso el valor num
	for _, num := range numeros {
		total += num
	}
	return total
}

func main() {
	result := sumar(5 + 5 + 5)
	fmt.Println(result)
}
