package main

import "fmt"

func main() {
	fmt.Print("Enter a Distance in Feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet * .3048
	fmt.Printf("%.2f converted to Meters: %.2f\n", feet, meters)
}
