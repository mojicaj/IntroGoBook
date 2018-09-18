package main

import (
	"fmt"
)

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

func makeOddGenerator() func() int {
	i := int(1)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
