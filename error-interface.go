package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	var contohError error = errors.New("Contoh error")
	fmt.Println(contohError.Error())

	result, err := pembagian(10, 0)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
