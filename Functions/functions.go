package main

import "fmt"

func add(a, b int) int {

	return a + b

}

func plusPlus(a string, b int, c int) int {

	fmt.Println("Message:", a)
	return b + c

}

func main() {

	fmt.Println("Sum1", plusPlus("Test Message", 13, 12))
	fmt.Println("Sum2:", add(12, 12))

}
