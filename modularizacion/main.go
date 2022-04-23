package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	// mensaje.Saluda()
	// mensaje.Print()

	// cuadrado := figuras.Cuadrado{Ancho: 2, Altura: 2.2}
	// circulo := figuras.Circulo{Radio: 5}

	// figuras.CalcularMedidas(&cuadrado)
	// figuras.CalcularMedidas(&circulo)

	// Encapsulamiento
	user1 := models.User{}
	user1.Constructor("Seba", 33)
	fmt.Println(user1)

	user2 := models.User{}
	user2.SetName("Seba 2")
	user2.SetEdad(333)

	fmt.Println("Nombre: ", user2.GetName())
	fmt.Println("Edad: ", user2.GetEdad())

}
