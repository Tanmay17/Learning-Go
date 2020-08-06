package main

import (
	"fmt"
)

const piVal = 3.14

func main() {

	fmt.Println(piVal)

	radius := 55

	fmt.Println("Radius is", radius)

	fmt.Println("Circumference of Circle:", piVal*float32(radius))

	piVal := 3.143
	fmt.Println("Updated Pi Value:", piVal)

}
