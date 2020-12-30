package main

import "fmt"

func main() {

	// slincing

	str := []string{"you", "we", "i", "to", "and"}

	str = append(str, "from")

	fmt.Printf("%s\n", str)

	s := str[2:6]

	fmt.Println(s)

	fmt.Printf("%s\n", str)

}
