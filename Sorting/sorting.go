package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{7, 2, 4}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
