package figuras

import "fmt"

// No es neecsario que la interface sea pública
type figura interface {
	area() float32
	perimetro() float32
}

func CalcularArea(f figura) {
	fmt.Println("El Area es :", f.area())
}

func CalcularPerimetro(f figura) {
	fmt.Println("El perímetro es :", f.perimetro())
}

func CalcularMedidas(f figura) {
	fmt.Println(f)
	fmt.Println("El área es :", f.area())
	fmt.Println("El perímetro es :", f.perimetro())
}
