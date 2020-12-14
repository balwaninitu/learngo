package main

import "fmt"

const (
	winter = 1

	summer = 3

	yearly = winter + summer
)

func main() {

	books := [yearly]string{
		"Home alone",
		"over the moon",

		"let it be",

		"dancing in the dark",
	}
	fmt.Printf(" \nBooks : %#v\n", books)

}
