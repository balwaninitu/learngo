package main

import (
	"fmt"
	"strings"
)

func main() {

	speed := 100

	fmt.Println("speed limit?", speed == 40 || speed == 50)
	//----------------------------------------------------

	fmt.Println(strings.Repeat("-", 30))

	color := "red"

	fmt.Println("is color red?", color == "green" || color == "dark green")

}
