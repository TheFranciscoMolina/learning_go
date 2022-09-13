package main

import (
	"fmt"
	"strings"
)

func reeplace(palabra string) {
	palabra = strings.Replace(palabra, "a", "i", 4)
	fmt.Println(palabra)
}

func main() {
	reeplace("no me la contes")
}
