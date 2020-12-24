package main

import "fmt"

func main() {
	names := [3]string{
		"Einstein", "Shepard",
		"Tesla",
	}

	books := [5]string{
		"Kafka's Revenge",
		"Stay Golden",
		"",
		"",
		"",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
