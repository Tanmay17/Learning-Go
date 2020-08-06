package main

import "fmt"

func main() {

	i := 1

	for i < 4 {
		fmt.Println(i)
		i++
	}

	for i := 4; i <= 9; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite Loop if No break")
		break
	}

	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
