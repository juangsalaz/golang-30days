package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Lamongan", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.City = "Surabaya"

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Malang"
	fmt.Println(address4)

	var alamat = Address{
		City:    "Subang",
		Provice: "Jawa Barat",
		Country: "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
