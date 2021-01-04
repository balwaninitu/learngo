package main

import "fmt"

func main() {

	weights := [...]float64{23, 24, 25, 26, 56}

	var sum float64

	// for _, w := range weights {

	// 	sum += w

	// }

	// fmt.Println(sum)

	// fmt.Printf("Average : %.2f\n", sum/float64(len(weights)))

	for i := 0; i < len(weights); i++ {

		sum += weights[i]

		fmt.Printf("sum %.2f\n", sum)

	}

	fmt.Printf("Average : %.2f\n", sum/float64(len(weights)))

}
