package main

import (
	"fmt"
	"os"
)

func main() {
	// hay que ejecutar el main.go desde la carpeta para poder ubicar el archivo

	// defer y recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups, al parecer el programa no finalizó de forma correcta")
			fmt.Println(error)
		}
	}()

	if file, error := os.Open("holas.txt"); error != nil {
		panic("No fue posible leer el archivo")
	} else {
		defer func() {
			file.Close()
			fmt.Println("Archivo cerrado")
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])

		fmt.Println(contenidoArchivo)
	}

}
