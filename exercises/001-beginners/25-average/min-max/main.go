package main

import "fmt"

func main() {

	s2 := [][]int{
		{86, 80, 75, 82, 80, 70},
		{40, 55, 60, 48, 50},
		{90, 85, 30, 86, 78, 75, 54},
	}

	max := s2[0][0]

	min := s2[0][0]

	for i := 0; i < len(s2); i++ {

		for j := 0; j < len(s2[i]); j++ {

			if max < s2[i][j] {

				max = s2[i][j]
			}

			if min > s2[i][j] {

				min = s2[i][j]
			}
		}

	}
	fmt.Printf("The Highest mark is %d\n", max)

	fmt.Printf("The lowest mark is %d\n", min)

}
