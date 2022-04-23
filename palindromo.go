package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hola Mundo"
	fmt.Printf("%s: %t \n", text, isPalindrome(text))

	text = "OSO"
	fmt.Printf("%s: %t \n", text, isPalindrome(text))
}

func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")
	text2 := reverse(text)

	if text == text2 {
		return true
	}
	return false
}

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInvertido := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertido = append(arrayInvertido, arrayCadena[i])
	}

	return strings.Join(arrayInvertido, "")
}
