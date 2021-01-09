package main

import "fmt"

type sum func(int, int) int

type item struct {
	grocery string

	vegetables string

	gprice int

	vprice int

	sum
}

func main() {

	items := item{

		grocery:    "dal",
		vegetables: "onion",

		gprice: 45,

		vprice: 56,

		sum: func(gprice int, vprice int) int {

			return gprice + vprice
		},
	}

	fmt.Println("sum", items.sum(items.gprice, items.vprice))

}
