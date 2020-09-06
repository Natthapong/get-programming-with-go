package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id
}

func main() {
	c := make(chan int)
	count := 5
	for i := 0; i < count; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < count; i++ {
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
