package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)

	mid := len(items) / 2

	mid1 := items[mid-1 : mid+2]

	sort.Strings(mid1)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
