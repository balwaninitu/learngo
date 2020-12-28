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
		locs                   []string
		size, bed, bath, price []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {

		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		size = append(size, n)

		n, _ = strconv.Atoi(cols[2])
		bed = append(bed, n)

		n, _ = strconv.Atoi(cols[3])
		bath = append(bath, n)

		n, _ = strconv.Atoi(cols[4])
		price = append(price, n)

	}

	for _, h := range strings.Split(header, separator) {

		fmt.Printf("%-15s", h)
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {

		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", size[i])
		fmt.Printf("%-15d", bed[i])
		fmt.Printf("%-15d", bath[i])
		fmt.Printf("%-15d", price[i])
		fmt.Println()
	}

	var sum int

	for _, s := range size {

		sum += s

	}

	fmt.Printf("Average: %.2f\n", float64(sum)/float64(len(size)))

	sum = 0
	for _, b := range bed {

		sum += b
	}

	fmt.Printf("Average: %.2f\n", float64(sum)/float64(len(bed)))

	sum = 0

	for _, ba := range bath {

		sum += ba

	}

	fmt.Printf("Average: %.2f\n", float64(sum)/float64(len(bath)))

	sum = 0

	for _, p := range price {

		sum += p

	}

	fmt.Printf("Average: %.2f\n", float64(sum)/float64(len(price)))

}
