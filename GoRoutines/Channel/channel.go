package main

import "fmt"

func main() {

	msg := make(chan string)

	go func() {
		msg <- "ping"
	}()

	msgs := <-msg
	fmt.Println(msgs)

}
