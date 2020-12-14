package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	const (
		usage               = " Usage :  [username]  [password]"
		username, username2 = "jack", "inac"
		password, password2 = "1888", "1879"
		errUser             = "Access denied for to %q.\n"
		errPwd              = "Invalid password for %q.\n"
		accessOK            = "Access granted to %q.\n"
	)

	if len(args) != 3 {

		fmt.Println(usage)

		return
	}

	u, p := args[1], args[2]

	if u != username && u != username2 {

		fmt.Printf(errUser, u)

	} else if u == username && p == password ||
		u == username2 && p == password2 {

		fmt.Printf(accessOK, u)

	} else {

		fmt.Printf(errPwd, u)
	}

}
