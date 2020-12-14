package main

import "fmt"

func main() {

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.87
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v miilions km\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	fmt.Printf("%v is away %v\n", planet, distance)
	fmt.Printf("%v is away %v\n Think %[1]v is %[2]v distnace away OMG! ", planet, distance)

}
