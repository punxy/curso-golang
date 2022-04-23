package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	channels := make(chan string)

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, srv := range servidores {
		// revisarServidor(srv) // SIN CONCURRENCIA
		// Para indicar que se ejecute en multiples hilos se agrega la palabra reservada GO
		go revisarServidor(srv, channels)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-channels)
	}

	tiempoEjecucion := time.Since(inicio)

	fmt.Println("Tiempo ejecución: ", tiempoEjecucion.Seconds())
}

func revisarServidor(servidor string, channel chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		// fmt.Printf("%s\t: No está disponible el servidor\n", servidor)
		channel <- servidor + " : No está disponible"

	} else {
		// fmt.Printf("%s\t: Está funcionando\n", servidor)
		channel <- servidor + " : Sí está disponible"

	}

}
