package main

import (
	"errors"
	"fmt"
)

func main() {

	result, error := dividir(5, 0)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}

func dividir(dividendo, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("No es posible dividir por cero")
	}

	result := dividendo / divisor
	return result, nil
}
