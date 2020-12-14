package main

import "fmt"

func main() {

	var (
		vehicle = "Car"
		speed   = 500
		mileage = 89.4
		color   = "red"
		isSUV   = true
	)

	fmt.Printf("I have a %v\nA %v color %[1]v\n", vehicle, color)
	fmt.Printf("Its mileage is %v & speed is  %vkm/hr\nYou cant beat such %[1]v & %[2]v\n", mileage, speed)
	fmt.Printf("It is %v its SUV\n", isSUV)
	fmt.Printf("I Love my %v & its %v color too\n", vehicle, color)

}
