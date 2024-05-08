package main

import "fmt"

// nolint: revive // DO NOT delete this comment!
func main() {
	// DO NOT delete this line! It takes the input numbers and generates the 3x3 matrix `A`.
	var A = generateMatrix()

	// Declare and initialize the transposed version of the 3x3 matrix `A` below:
	var transposedA [3][3]int

	for key, _ := range A {
		for coll, val := range A[key] {
			transposedA[coll][key] = val
		}
	}

	// Print each row of the 3x3 `transposedA` matrix below:
	fmt.Println(transposedA[0])
	fmt.Println(transposedA[1])
	fmt.Println(transposedA[2])
}

func generateMatrix() [3][3]int {
	return [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
}
