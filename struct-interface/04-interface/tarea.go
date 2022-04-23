package main

import (
	"fmt"
	"math"
)

type Figura interface {
	area() float32
	perimetro() float32
}

type Cuadrado struct {
	ancho  float32
	altura float32
}

type Circulo struct {
	radio float32
}

func (c *Cuadrado) area() float32 {
	return c.ancho * c.altura
}

func (c *Cuadrado) perimetro() float32 {
	return (2 * c.ancho) + (2 * c.altura)
}

func (c *Circulo) area() float32 {
	return (math.Pi * (c.radio * c.radio))
}

func (c *Circulo) perimetro() float32 {
	return 2 * math.Pi * c.radio
}

func calcularArea(f Figura) {
	fmt.Println("El Area es :", f.area())
}

func calcularPerimetro(f Figura) {
	fmt.Println("El perímetro es :", f.perimetro())
}

func calcularMedidas(f Figura) {
	fmt.Println(f)
	fmt.Println("El área es :", f.area())
	fmt.Println("El perímetro es :", f.perimetro())

}

func main() {

	cuadrado := Cuadrado{2, 2.2}
	circulo := Circulo{5}

	// calcularArea(&cuadrado)
	// calcularArea(&circulo)
	// calcularPerimetro(&cuadrado)
	// calcularPerimetro(&circulo)

	calcularMedidas(&cuadrado)
	calcularMedidas(&circulo)

}
