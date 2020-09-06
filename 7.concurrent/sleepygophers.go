package main

import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}

func main() {
	count := 5
	for i := 0; i < count; i++ {
		go sleepyGopher()
	}
	time.Sleep(4 * time.Second)
}
