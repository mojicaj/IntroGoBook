package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter an integer: ")
	var i int
	fmt.Scanf("%v", &i)
	fmt.Println(half(i))
}

func half(i int) (int, bool) {
	r := i / 2
	if i%2 == 0 {
		return r, true
	}
	return r, false
}
