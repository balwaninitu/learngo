package main

import "fmt"

const (
	winter = 1

	summer = 3

	yearly = winter + summer
)

func main() {

	var books [yearly]string

	books[0] = "home alone"

	books[1] = "bees home"

	books[2] = "over the moon"

	books[3] = "dancing in the dark"

	var (
		wBooks [winter]string

		sbooks [summer]string
	)

	wBooks[0] = books[0]

	// sbooks[0] = books[1]

	// sbooks[1] = books[2]

	// sbooks[2] = books[3]

	// for i := 0; i < len(sbooks); i++ {

	// 	sbooks[i] = books[i+1]

	for i := range sbooks {

		sbooks[i] = books[i+1]
	}

	var published [len(books)]bool

	published[0] = true

	published[len(books)-1] = true

	fmt.Println("/nPublished books:")

	for i, ok := range published {

		if ok {

			fmt.Printf("+  %s\n", books[i])
		}
	}

	// fmt.Printf("\nwbooks : %#v\n", wBooks)

	// fmt.Printf("\nsbooks : %#v\n", sbooks)

}
