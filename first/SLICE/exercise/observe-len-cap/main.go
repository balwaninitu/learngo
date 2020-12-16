package main

import "fmt"

func main() {

	games := []string{}

	fmt.Println("Length", len(games))

	fmt.Println("capacity", cap(games))

	games = append(games, "pacman", "mario", "tetris", "doom")

}
