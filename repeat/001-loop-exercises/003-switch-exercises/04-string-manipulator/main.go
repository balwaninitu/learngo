package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {

		fmt.Printf(" [command] [string]\nAvailable commands: lower, upper and title\n")

		return
	}

	cmd, str := os.Args[1], os.Args[2]

	var res string

	switch {

	case cmd == "lower":

		res = strings.ToLower(str)

	case cmd == "upper":

		res = strings.ToUpper(str)

	case cmd == "title":

		res = strings.ToTitle(str)

	default:

		res = "unknow command"

	}

	fmt.Printf("%s %s", cmd, res)

}
