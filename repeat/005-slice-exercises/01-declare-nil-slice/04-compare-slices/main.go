package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ", ")

	sort.Strings(namesC)

	sort.Strings(namesB)

	fmt.Println(namesC)
	fmt.Println(namesB)

	if len(namesC) == len(namesB) {

		for i := range namesC {

			if namesC[i] == namesB[i] {

				fmt.Println("they are equal")

				return

			}

		}

	}

	fmt.Println("they are not equal")
}
