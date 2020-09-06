package main

import (
	"fmt"
	"strings"
)

func sourceGopher(values []string, downstream chan string) {
	for _, v := range values {
		downstream <- v
	}
	downstream <- ""
}

func filterGopher(upstream, downstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			downstream <- ""
			return
		}
		if !strings.Contains(v, "bad") {
			downstream <- v
		}
	}
}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	values := []string{"hello world", "a bad apple", "goodbye all"}
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(values, c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
