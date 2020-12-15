package main

import "fmt"

func main() {

	var todo []string

	todo = append(todo, "sing", "code", "play", "dance")

	tomorrow := []string{"learngo", "play with Riddham"}

	todo = append(todo, tomorrow...)

	fmt.Printf(" %T %s %d", todo, todo, len(todo))

	fmt.Println(todo[3])
}
