package main

import "fmt"

func main() {

	var fruts map[string]int
	fruts = make(map[string]int)
	fruts["banana"] = 2
	fruts["pineaple"] = 5
	fruts["aple"] = 3

	var vegetable = make(map[string]bool)
	vegetable["tomato"] = true
	vegetable["brocolis"] = true
	vegetable["wathermelon"] = false

	type Person struct {
		Sufix    string
		LastName string
		Income   float32
	}

	names := map[string]Person{}
	names["maria"] = Person{"Ms", "Maria", "Something", 134000.34}
	names["Jhon"] = Person{"Sr", "Jhon", "Somethinelse", 65000.32}

	for key, val := range fruts {
		fmt.Printf("You have %d %s\n", val, key)
	}

	for key, val := range vegetable {

		awnser := "No"
		if val {
			awnser = "Yeess"
		}

		fmt.Printf("Is %s a vegetable: %s\n", key, awnser)
	}

	for _, val := range names {
		fmt.Printf("%s %s get £%.2f per anual \n", val.Sufix, val.LastName, val.Income)
	}

}
