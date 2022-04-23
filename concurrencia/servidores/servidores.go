package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, srv := range servidores {

		// SIN CONCURRENCIA
		revisarServidor(srv)
	}

	tiempoEjecucion := time.Since(inicio)

	fmt.Println("Tiempo ejecución: ", tiempoEjecucion)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err == nil {
		fmt.Printf("%s\t: No está disponible el servidor\n", servidor)
	} else {
		fmt.Printf("%s\t: Está funcionando\n", servidor)
	}

}
