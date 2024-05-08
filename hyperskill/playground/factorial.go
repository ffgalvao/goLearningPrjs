package main

/*
Factorial https://hyperskill.org/learn/step/22041
You need to create an anonymous function that calculates the factorial of a number and sets the value of the factorial variable.

Take notice that the anonymous function is called directly, and the result is saved to the factorial variable.
The factorial of a number is the result of multiplying integer numbers from 1 to the number. For example 5! = 1 * 2 * 3 * 4 * 5 = 120.

Sample Input 1:
1
Sample Output 1:
1

*/

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	factorial := func() uint64 {
		var f uint64 = 1
		for i := uint64(1); i <= uint64(number); i++ {
			f *= i
		}

		return f
	}()

	fmt.Println(factorial)
}
