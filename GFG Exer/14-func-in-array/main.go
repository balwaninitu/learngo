package main

import "fmt"

func myfun(a [9]int, size int) int {

	var i, sum, r int

	for i = 0; i < size; i++ {

		sum += a[i]
	}

	r = sum / size

	return r
}

func main() {

	var arr = [9]int{67, 59, 29, 35, 4, 34, 56, 78, 34}

	var res int

	res = myfun(arr, 9)

	fmt.Printf("Average %d", res)

}
