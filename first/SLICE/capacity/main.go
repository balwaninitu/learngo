package main

import "fmt"

func main() {

	var games []string

	games = []string{"marion", "ludo", "tetris"}

	part := games

	fmt.Println(games)

	fmt.Println(part[0:2])

	fmt.Println(part[:0])

	fmt.Println(part[:cap(part)])

}
