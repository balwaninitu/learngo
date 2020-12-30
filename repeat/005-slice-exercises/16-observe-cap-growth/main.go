package main

import "fmt"

func main() {
	var (
		num []int

		oldCap float64
	)

	for len(num) < 10 {

		c := float64(cap(num))

		if c == 0 || c != oldCap {

			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(num), c, c/oldCap)

		}

		oldCap = c

		num = append(num, 1)

	}

}
