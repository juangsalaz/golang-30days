package main

import "fmt"

func Random() interface{} {
	return true
}

func main() {
	random := Random()
	// stringRandom := random.(string)
	// fmt.Println(stringRandom)

	// intRandom := random.(int)
	// fmt.Println(intRandom)

	switch value := random.(type) {
	case string:
		fmt.Println("String", value, "is String")
	case int:
		fmt.Println("Int", value, "is Int")
	default:
		fmt.Println("Unknown type")
	}
}
