package main

import "fmt"

func main() {

	msg := make(chan string, 2)

	msg <- "Hello"
	msg <- "Duniya"

	fmt.Println(<-msg)
	fmt.Println(<-msg)

}
