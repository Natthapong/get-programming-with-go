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
	go sleepyGopher()
	time.Sleep(4 * time.Second)
}
