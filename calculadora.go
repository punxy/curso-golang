package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var expresion string
	var result int

	fmt.Print("=> ")
	fmt.Scanln(&expresion)

	result = sumar(expresion)
	fmt.Printf("=> %d \n", result)
}

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	/*
	 * range devuelve el indice y el valor
	 * se pone un guión bajo para indicar que no lo usaremos
	 */
	for _, valor := range valores {
		// el segundo valor es un nil, que no usaremos
		num, error := strconv.Atoi(valor)

		if error != nil {
			fmt.Println(error)
		} else {
			result += num
		}
	}

	return result
}
