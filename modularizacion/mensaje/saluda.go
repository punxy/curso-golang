package mensaje

import "fmt"

// minúscula para privada
const mensaje = "mensaje desde constante privada"

// Con Mayúscula para que el método sea público
func Saluda() {
	fmt.Println("Hola desde paquete mensaje, archivo Saluda")
}

// Con Minúscula para que el método sea privado
func chao() {
	fmt.Println("Chao desde paquete mensaje, archivo Saluda")
}

func Print() {
	// imrpimimos una constante privada
	fmt.Println(mensaje)
}
