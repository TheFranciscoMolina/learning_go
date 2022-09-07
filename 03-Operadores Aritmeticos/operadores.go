package main

import "fmt"

func main (){

	var result int
	a := 200
	b := 100
	//Suma
	result = a+b
	fmt.Println("Suma:", result)
	//Resta (solo se usa el igual porque ya fue declarado)
	result = a-b
	fmt.Println( "Resta:", result)
	result = a*b
	fmt.Println("Multiplicacion",result)
	result = a / b
	fmt.Println( "Division: ",result)

	var divi float64 = 3.0 / 2.5
	fmt.Println( "Division decimales:",divi)

	result = a % b
	fmt.Println("Modulo es", result)


}