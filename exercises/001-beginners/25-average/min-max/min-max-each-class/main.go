package main

import "fmt"

func main() {

	s2 := [][]int{
		{86, 80, 75, 82, 80, 70},
		{40, 55, 60, 48, 50},
		{90, 85, 30, 86, 78, 75, 54},
	}

	var eachMax, eachMin int

	for i := 0; i < len(s2); i++ {

		max := s2[i][0]

		min := s2[i][0]

		for j := 0; j < len(s2[i]); j++ {

			if s2[i][j] > max {

				max = s2[i][j]
			}

			if s2[i][j] < min {

				min = s2[i][j]
			}

			eachMax = max

			eachMin = min

		}

		fmt.Printf("The highest mark of class %d is %d\n", i, eachMax)

		fmt.Printf("The lowest mark of class is %d %d\n", i, eachMin)

	}

}
