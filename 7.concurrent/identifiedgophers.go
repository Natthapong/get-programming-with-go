package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, "snore ...")
}
func main() {
	count := 5
	for i := 0; i < count; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
