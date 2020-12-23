package main

import (
	"fmt"
	"os"
)

const (
	user1, user2 = "riddham", "nitu"
	p1, p2       = "2012", "1984"
)

func main() {

	if len(os.Args) != 3 {

		fmt.Println("Uage: [Username] [Password]")

		return
	}

	u := (os.Args[1])

	p := (os.Args[2])

	var res string
	switch {

	case u == user1 && p == p1 || u == user2 && p == p2:

		res = "AccessOk"

	case u != user1 && u != user2:

		res = "Access denied"

	default:

		res = "Invalid password"

	}

	fmt.Printf("%s for %s\n", res, u)

}
