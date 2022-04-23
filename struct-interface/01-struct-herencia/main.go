package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) saluda() {
	fmt.Printf("Hola %s, tienes %d \n", p.nombre, p.edad)
}

func (p *Persona) setNombre(nombre string) {
	p.nombre = nombre
}

type Empleado struct {
	Persona
	sueldo int
}

func main() {
	fmt.Println("\n\n*** Struct / Clase ***\n ")

	persona01 := Persona{"Seba", 33}
	persona01.saluda()

	persona02 := Persona{
		nombre: "Seba2",
		edad:   33,
	}

	persona02.saluda()

	// Herencia
	fmt.Println("\n*** Herencia ***\n ")
	em := Empleado{sueldo: 2300}
	em.setNombre("Pedro")
	em.edad = 25
	em.saluda()
	fmt.Println(em)
	fmt.Println(em.Persona)
}
