package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fac := n * factorial(n-1)
	return fac
}

func main() {
	fmt.Println(factorial(6))
}
