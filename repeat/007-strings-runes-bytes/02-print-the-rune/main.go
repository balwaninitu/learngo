package main

import "fmt"

func main() {
	const word = "console"

	for _, r := range word {

		fmt.Printf("%c", r)

		fmt.Printf("\tdecimal : %[1]d\n", r)
		fmt.Printf("\thex : %[1]x\n", r)
		fmt.Printf("\tbin : %[1]b\n", r)
	}

	var b []byte

	b = []byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}

	fmt.Printf("%s", b)

}
