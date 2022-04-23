package main

import "fmt"

type User struct {
	name   string
	mail   string
	active bool
}

type Student struct {
	user User
	code string
}

// Relación de UNO a MUCHO
type Curse struct {
	name   string
	code   string
	videos []Video
}

type Video struct {
	name   string
	number int
	curse  Curse
}

func main() {

	// Relación UNO a UNO
	usr1 := User{
		name:   "Seba",
		mail:   "sebasilvac88@gmail.com",
		active: true,
	}

	student1 := Student{
		user: usr1,
		code: "code_009988",
	}

	fmt.Println(usr1, student1)

	v1 := Video{name: "Video1", number: 1}
	v2 := Video{name: "Video2", number: 2}

	curso := Curse{
		code:   "code_curso_123",
		name:   "Curso de Go",
		videos: []Video{v1, v2},
	}

	v1.curse = curso
	v2.curse = curso

	fmt.Println("Nombre: ", v1.curse.name)
	fmt.Println("Código: ", v2.curse.code)

	for _, video := range curso.videos {
		fmt.Println("Nombre Video: ", video.name)
	}

}
