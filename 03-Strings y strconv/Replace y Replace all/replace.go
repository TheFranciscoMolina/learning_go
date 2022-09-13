package main

import (
	"fmt"
	"strings"
)

func toReplace(palabra string) {
	palabra = strings.Replace(palabra, "a", "i", 4)
	fmt.Println(palabra)
}

func main() {
	toReplace("no me la contes")
}
