package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func main() {
	address1 := Address{"Lamongan", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.City = "Surabaya"

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
