package main

import "fmt"

func main() {
	fruit := "banana"

	switch fruit {
	case "apple":
		fmt.Println("is an apple")
	default:
		fmt.Println("is not an apple")
	}

	defer fmt.Println("world")

	fmt.Println("hello")
}
