package main

import "fmt"

func main() {

	const home = "Milky way Galaxy"

	const length = len(home)

	fmt.Printf("There are %d characters inside %q.\n", length, home)

}
