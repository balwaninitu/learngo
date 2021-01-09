package main

import "fmt"

func main() {

	arr := [...][3]int{

		{3, 4, 5},
		{6, 5, 1},
		{9, 6, 1},
	}

	var i, j int

	for i = 0; i < 3; i++ {

		for j = 0; j < 3; j++ {

			fmt.Println(arr[i][j])

		}

	}

	fmt.Println(arr)

}
