package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)

	fmt.Println("len:", len(s))

	s = append(s, "d", "e")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	fmt.Println("sl1", c[2:5])

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
