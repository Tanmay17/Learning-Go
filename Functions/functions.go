package main

import "fmt"

func add(a, b int) int {

	return a + b

}

func multipleReturn(a string, b int, c int) (int, string) {

	return b + c, a

}

func main() {

	sum, msg := multipleReturn("Test Message", 13, 12)
	fmt.Println("Sum1:", sum)
	fmt.Println("Msg:", msg)
	fmt.Println("Sum2:", add(12, 12))

}
