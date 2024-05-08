package main

/*
Count odd numbers
Here is the file numbers.txt; it contains a sequence of integers.
Each number is in a new line.

Read the file numbers.txt, then use a for scanner.Scan() loop to iterate over each
line of the file, count the number of odd numbers in the oddNumberCount variable,
and finally print the total amount of odd numbers the file has.
*/

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	filePath, _ := filepath.Abs("files/numbers.txt")
	file, err := os.Open(filePath) // open 'numbers.txt' here with the os.Open function
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var oddNumberCount int // oddNumberCount will be used to count the amount of odd numbers in the file

	scanner := bufio.NewScanner(file) // create a new scanner for the file

	// Use the for scanner.Scan() loop to read the file line by line
	for scanner.Scan() {
		// DO NOT delete the code block below, it parses the read number into an integer:
		var number int
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		// Write the code that checks if 'number' is odd and then increments 'oddNumberCount' below:
		if number%2 != 0 {
			oddNumberCount++
		}
	}
	fmt.Println(oddNumberCount) // print the total oddNumberCount here!
}
