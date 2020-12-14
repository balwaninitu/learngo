package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("[your name]")

		return

	}

	name := args[0]

	feelings := [...]string{

		"  sad :(",

		"  terrible :(",

		"  happy :)",

		"  awesome :)",

		"  good :)",
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(feelings))

	fmt.Printf("%s feels %s\n", name, feelings[n])

}
