package main

import "fmt"

func main() {
	dwarf := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarf)

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(len(planets))

	for i := 0; i < len(planets); i++ {
		fmt.Println(i, planets[i])
	}

	for k, v := range planets {
		fmt.Println(k, v)
	}
}
