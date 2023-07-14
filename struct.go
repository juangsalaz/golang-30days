package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var juang Customer
	juang.Name = "Juang"
	juang.Address = "Lamongan"
	juang.Age = 30

	// fmt.Println(juang)

	juang.sayHello("prabowo")
}
