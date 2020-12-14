package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Riddham           is a nice girl            "

	name = strings.TrimRight(name, "  is ")

	fmt.Println(name)

}
