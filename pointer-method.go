package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	juang := Man{"Juang"}
	juang.Married()

	fmt.Println(juang.Name)
}
