package main

import "fmt"

func main() {

	servers := [4]bool{

		false, true, false, false,
	}

	fmt.Printf("%#v\n", servers)

	fmt.Printf("%T\n", servers)

	fmt.Printf("%t\n", servers)

}
