package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(7)

	fmt.Printf("%d\n", num)

}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func random(min int, max int) int {
// 	return rand.Intn(max-min) + min
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	randomNum := random(1, 7)
// 	fmt.Printf("Random number: %d\n", randomNum)
// }
