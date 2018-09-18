package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	fmt.Print("Enter an integer for x: ")
	fmt.Scanf("%v", &x)
	fmt.Print("Enter an integer for y: ")
	fmt.Scanf("%v", &y)
	swap(&x, &y)
	fmt.Printf("After swapping: x=%v y=%v\n", x, y)
}

func swap(x, y *int) {
	n := *x
	*x = *y
	*y = n
}
