package main

import "fmt"

func main() {

	/* 	Ejm con descuentos:

	   	Descuento del 10% para compras de hasta $100
	   	Descuento del 20% para combras de hasta $200
	   	El iva es del: 21% */

	var total, descuento float64
	var datoDeDescuento string

	fmt.Print("***Consumidor final***")
	fmt.Print("Total: ")
	fmt.Scanln(&total)

	if total >= 0 {
		if total > 100 {
			datoDeDescuento = "10%"
			descuento = total * 0.10
			fmt.Printf("Wow! Tienes un %d de descuento!", datoDeDescuento)
			fmt.Printf("Equivale a: %d", descuento)
		}
		if total > 200 {
			datoDeDescuento = "20%"
			descuento = total * 0.20
			fmt.Printf("Wow! Tienes un %d de descuento!", datoDeDescuento)
			fmt.Printf("Equivale a: %d", descuento)
		}
	} else {
		fmt.Println("Error: El monto total no puede ser menor a 0.")
	}
}
