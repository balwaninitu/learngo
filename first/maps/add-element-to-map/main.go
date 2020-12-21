package main

import "fmt"

func main() {

	phones := map[string]string{

		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	//	product := map[int]bool{

	//		617841573 : true,
	//		879401371 : false,
	//          576872813 : true,

	//	}

	//	multiphone := map[string][]string{

	//		"bowen":  {"202-555-0179"},
	//		"dulin": {"03.37.77.63.06"},
	//		"greco": {"03489940240"},

	//	}

	//	basket := map[int]map[int]int{

	//		100 : {67542312 : 4, 45890712: 2},
	//		101 : {87512310 : 4, 65860714: 2},
	//		102 : {67542312 : 4, 61897715: 2},

	//	}

	who, phone := "dulin", "N/A"
	if v, ok := phones[who]; ok {
		phone = v
	}
	fmt.Printf("%s's phone number: %s\n", who, phone)

}
