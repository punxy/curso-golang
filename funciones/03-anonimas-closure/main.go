package main

import (
	"fmt"
	"strings"
)

func main() {

	// función anónima
	func() {
		fmt.Println("Hola Mundo, desde función anónima")
	}() // le agregamos los parenyecis para que se ejecute automáticamente

	anom := func(nombre string) string {
		return fmt.Sprintf("Hola %s !!!", nombre)
	}

	fmt.Println(anom("Seba"))

	funRepeat3 := repeat(3)

	fmt.Println(funRepeat3("Hola"))
}

// Closures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}
