package main

import "fmt"

func main() {
	mensaje, result := sumar("Seba", 1, 2, 3, 4, 5)
	fmt.Println(mensaje, result)
}

// una función variadica recibe y retornar argumentos indefinidos
func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es:", nombre)

	total := 0
	for _, num := range numeros {
		total += num
	}

	return mensaje, total
}
