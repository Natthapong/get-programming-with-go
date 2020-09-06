package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
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
		fmt.Println("gopher", gopherID, "has finished sleeping")
	}
}
