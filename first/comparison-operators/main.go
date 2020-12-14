package main

import (
	"fmt"
	"strings"
)

func main() {

	speed := 100

	fast := speed >= 80

	slow := speed < 20

	fmt.Printf("speed is fast? %t.\n", fast)

	fmt.Printf("speed is slow? %t.\n", slow)

	fmt.Printf("fast type is %T.\n", fast)

	fmt.Printf("is it 100 mph? %t.\n", speed == 100)

	fmt.Printf("is it not 100 mph %t \n", speed != 100)

	fmt.Println(strings.Repeat("-", 25))

	//----------------------------------------------

	speedB := 150.5

	faster := speedB > 100

	fmt.Println("is other care going faster?", faster)

	faster = speedB > float64(speed)

	fmt.Println(faster)

}
