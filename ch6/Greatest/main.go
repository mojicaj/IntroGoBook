package main

import (
	"fmt"
)

func main() {

	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	fmt.Println("Greates number is ", greatest(x...))
}

func greatest(nums ...int) int {
	greater := nums[0]

	for _, v := range nums {
		if v > greater {
			greater = v
		}
	}
	return greater
}
