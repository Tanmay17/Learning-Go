package main

import "fmt"

func main() {

	a := make(map[string]int)

	a["key1"] = 1
	a["key2"] = 2
	a["key3"] = 3

	fmt.Println("map:", a)

	v1 := a["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(a))

	delete(a, "key2")
	fmt.Println("map:", a)

	_, prs := a["key2"]
	fmt.Println("prs:", prs)

	v := map[string]int{"foo": 112, "bar": 232}
	fmt.Println(v)

}
