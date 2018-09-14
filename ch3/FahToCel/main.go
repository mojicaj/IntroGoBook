package main

import (
	"fmt"
)

// convert a temperature from Fahrenheit to Celsius
func main() {
	fmt.Print("Enter a Fahrenheit temperature: ")
	var fah float64
	fmt.Scanf("%f", &fah)

	cel := (fah - 32) * 5 / 9
	fmt.Printf("%.2f converted to Celsius: %.2f\n", fah, cel)
}
