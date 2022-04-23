package models

type User struct {
	name string
	edad int
}

func (u *User) Constructor(name string, edad int) {
	u.name = name
	u.edad = edad
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEdad(edad int) {
	u.edad = edad
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEdad() int {
	return u.edad
}
