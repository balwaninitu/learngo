package main

import (
	"fmt"
	"strings"
)

func main() {

	books := [...]string{"over the moon",

		"let the party on", "night is dark", "its not over yet",
		"howdy modi"}

	upper, lower := books, books

	for i := range books {

		upper[i] = strings.ToUpper(upper[i])

		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("books : %q\n", books)

	fmt.Printf("upper : %q\n", upper)

	fmt.Printf("upper : %q\n", lower)
}
