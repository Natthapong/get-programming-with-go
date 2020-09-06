package main

import (
	"fmt"
	"strings"
)

func sourceGopher(strings []string, downstream chan string) {
	for _, str := range strings {
		downstream <- str
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for str := range upstream {
		if !strings.Contains(str, "bad") {
			downstream <- str
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for str := range upstream {
		fmt.Println(str)
	}
}

func main() {
	strings := []string{"hello world", "a bad apple", "goodbye all"}
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(strings, c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
