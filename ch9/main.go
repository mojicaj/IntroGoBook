package main

import (
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}

	fmt.Println("Avg: ", average(xs))
	fmt.Println("Min: ", min(xs))
	fmt.Println("Max: ", max(xs))
}

// Finds the average of a series of numbers
func average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the lowest value of a series of numbers
func min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	lowest := xs[0]

	for _, v := range xs {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}

// Finds the highest value of a series of numbers
func max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	highest := xs[0]

	for _, v := range xs {
		if v > highest {
			highest = v
		}
	}
	return highest
}
