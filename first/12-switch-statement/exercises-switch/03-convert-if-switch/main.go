package main

import (
	"fmt"
	"os"
)

const (
	usage       = "Usage : [Username] [Password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Password invalid for %q.\n"
	accessOK    = "Access granted for %q.\n"
	user, user2 = "jack", "inac"
	pass, pass2 = "1888", "1879"
)

func main() {

	args := os.Args

	if len(args) != 3 {

		fmt.Println(usage)

		return

	}

	u, p := args[1], args[2]

	switch {

	case u != user && u != user2:

		fmt.Printf(errUser, u)

	case u == user && p == pass:

		fallthrough

	case u == user2 && p == pass2:

		fmt.Printf(accessOK, u)

	default:

		fmt.Printf(errPwd, u)

	}

}
