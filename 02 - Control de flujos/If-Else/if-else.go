package main

import "fmt"

func main() {

	/* 	Ejm con descuentos:

	   	Descuento del 10% para compras de hasta $100
	   	Descuento del 20% para combras de hasta $200
	   	El iva es del: 21% */

	var total, descuento float64
	var datoDeDescuento string

	fmt.Println("****Facturador:***")
	fmt.Scanln(&total)

	if total >= 0 {
		if total <= 100 {
			datoDeDescuento = "10%"
			descuento = total * 0.10
		} else if total > 100 && total <= 200 {
			datoDeDescuento = "20%"
			descuento = total * 0.20
		} else if total < 200 {
			datoDeDescuento = "30%"
			descuento = total * 0.30
		}
	} else {
		fmt.Println("Error: El monto total no puede ser menor a 0.")
	}

	//Operaciones:
	montoConDescuento := total - descuento
	iva := montoConDescuento * 0.25
	totalFinal := montoConDescuento + iva

	//Impromo los datos:

	fmt.Printf("WOW! Tienes un descuento del %s \n", datoDeDescuento)
	fmt.Println("El total es de: $", total)
	fmt.Println("Se desconto: $", montoConDescuento)
	fmt.Println("El IVA es de: $", iva)
	fmt.Println("El total a pagar es de: $", totalFinal)
	fmt.Println("**Cerrando sistema**")
}
