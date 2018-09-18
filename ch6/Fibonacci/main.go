package main

import "fmt"

func main() {
	fmt.Print("Enter an integer: ")
	var n int
	fmt.Scanf("%v", &n)
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
