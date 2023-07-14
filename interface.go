package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHi(name HasName) {
	fmt.Println("Hi", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	var juang Person
	juang.Name = "Juang"

	SayHi(juang)

	var gajah Animal
	gajah.Name = "Gajah"

	SayHi(gajah)

}
