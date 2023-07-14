package main

import "fmt"

func logging() {
	message := recover()
	fmt.Println("Error dengan message", message)
	fmt.Println("selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run appliction")

	result := 10 / value

	fmt.Println(result)
}

func main() {
	runApplication(0)
}
