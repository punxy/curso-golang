package main

import (
	"paquetes/figuras"
	"paquetes/mensaje"
)

func main() {
	mensaje.Saluda()
	mensaje.Print()

	cuadrado := figuras.Cuadrado{Ancho: 2, Altura: 2.2}
	circulo := figuras.Circulo{Radio: 5}

	figuras.CalcularMedidas(&cuadrado)
	figuras.CalcularMedidas(&circulo)
}
