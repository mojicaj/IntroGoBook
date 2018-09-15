package main

import (
	"fmt"
	"strconv"
)

func main() {
	var output string
	for i := 1; i <= 100; i++ {
		switch 0 {
		case i % 15:
			output = "FizzBuzz"
		case i % 3:
			output = "Fizz"
		case i % 5:
			output = "Buzz"
		default:
			output = strconv.Itoa(i)
		}
		fmt.Println(output)
	}
}
