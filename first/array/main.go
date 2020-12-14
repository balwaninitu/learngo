package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {

	var books [yearly]string

	books[0] = "Home alone"

	books[1] = "Home alone1"

	books[2] = "Home alone2"

	books[3] += books[0] + "3"

	fmt.Printf("Books :%#v\n", books)

}
