package main

import "fmt"

func main() {

	var (
		phones      map[string]string
		product     map[int]bool
		multiphones map[string][]string
		custID      map[int]map[int]int
	)

	fmt.Println(phones)
	fmt.Println(product)
	fmt.Println(multiphones)
	fmt.Println(custID)

	phones = map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240"}

	product = map[int]bool{

		617841573: true,
		879401371: false,
		576872813: true,
	}

	multiphones = map[string][]string{

		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06 03.37.70.50.05 02.20.40.10.04"},
		"greco": {"03489940240 03489900120"},
	}

	custID = map[int]map[int]int{

		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	who, phone := "Dulin", "N/A"

	if v, ok := phones[who]; ok {

		phone = v
	}

	fmt.Printf("%s's phone number : %s\n", who, phone)

	id, status := 879401371, "available"

	if !product[id] {

		status = "not" + status
	}

	whos, phone := "greco", "N/A"

	if phones := multiphones[whos]; len(phones) >= 2 {

		phone = phones[1]
	}

	fmt.Printf("%s's 2nd phone number: %s\n", who, phone)

	cid, pid := 101, 576872813
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, custID[cid][pid], pid)
}
