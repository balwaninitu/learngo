package main

import "fmt"

func main() {

	prev := [3]string{"home alone",

		"over the moon", "next to you",
	}

	var books [4]string

	for i, b := range prev {

		books[i] += b + "2nd ed."
	}

	books[3] = "witches"

	fmt.Printf("last year: \n%#v\n", prev)

	fmt.Printf("this year: \n%#v\n", books)

}
