package main

import "fmt"

type Perro struct{}
type Pez struct{}
type Pajaro struct{}

type Animal interface {
	mover() string
}

func (*Perro) mover() string {
	return "Soy un perro y puedo caminar"
}

func (*Pajaro) mover() string {
	return "Soy un pajaro y puedo volar"
}

func (*Pez) mover() string {
	return "Soy un pez y puedo nadar"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {
	perro := Perro{}
	pajaro := Pajaro{}
	pez := Pez{}

	moverAnimal(&perro)
	moverAnimal(&pajaro)
	moverAnimal(&pez)
}
