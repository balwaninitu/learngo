package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs []string

		Size, Beds, Baths, Price []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {

		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		Size = append(Size, n)

		n, _ = strconv.Atoi(cols[2])
		Beds = append(Beds, n)

		n, _ = strconv.Atoi(cols[3])
		Baths = append(Baths, n)

		n, _ = strconv.Atoi(cols[4])
		Price = append(Price, n)

	}

	for _, h := range strings.Split(header, separator) {

		fmt.Printf("%-15s", h)

	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {

		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", Size[i])
		fmt.Printf("%-15d", Beds[i])
		fmt.Printf("%-15d", Baths[i])
		fmt.Printf("%-15d", Price[i])
		fmt.Println()
	}

}
