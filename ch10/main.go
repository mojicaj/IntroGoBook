package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start: ", time.Now())
	sleep(time.Second * 5)
	fmt.Println("End: ", time.Now())
}

func sleep(d time.Duration) {
	select {
	case <-time.After(d):
		return
	}
}
