package main

import "fmt"

func main() {

	safe := true
	fmt.Println(safe)

	safe, age := false, 20

	fmt.Println(safe, age)

}
