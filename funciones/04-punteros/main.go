package main

import "fmt"

func main() {
	cadena := "Hola desde el main"
	fmt.Printf("%p \n", &cadena)

	fmt.Println("Antes de la función\n", cadena)

	//modificarValor(cadena)
	//fmt.Println("Después de la función normal", cadena)

	modificarConPuntero(&cadena)
	fmt.Println("Después de la función con puntero\n", cadena)

}

func modificarValor(cadena string) {
	fmt.Printf("%p \n", &cadena)
	cadena = "Hola desde función"
}

func modificarConPuntero(cadena *string) {

	// quitamos el % porque ya es una refreencia a la memoria
	fmt.Printf("%p \n", cadena)
	*cadena = "Hola desde función"
}
