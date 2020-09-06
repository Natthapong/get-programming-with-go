package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(4 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id
}

func main() {
	c := make(chan int)
	count := 5
	for i := 0; i < count; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(3 * time.Second)
	for i := 0; i < count; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}
