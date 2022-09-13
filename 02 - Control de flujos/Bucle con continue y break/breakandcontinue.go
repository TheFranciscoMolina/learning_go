package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {

		if i == 4 {
			fmt.Println("Continua")
			continue
		}
		fmt.Println("El valor es:", i)
		if i == 6 {
			break
		}

	}
}
