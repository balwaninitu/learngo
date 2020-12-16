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

	fmt.Println("Original:\n", items)

	mid := len(items) / 2

	fmt.Println("mid", mid)

	smid := items[mid-1 : mid+2]

	fmt.Println("smid:", smid)

	sort.Strings(smid)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
