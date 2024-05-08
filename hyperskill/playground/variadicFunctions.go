package main

/*
Greeting code https://hyperskill.org/learn/step/21807
Here is a function description:

func Greeting(prefix string, name ...string) []string {...}
You need to write the code of a function that returns a slice of strings. Every string should have the following structure: prefix + whitespace + name.

For example:

Greeting("Hello,", "Rick", "Morty") // [Hello, Rick Hello, Morty]
In the input, you get three strings. The first string is a prefix, the other two are names. You need to output the result that the function returns line by line.

Sample Input 1:
Hello, Jerry Beth
Sample Output 1:
Hello, Jerry
Hello, Beth
*/

import "fmt"

func main1() {
	var prefix, name1, name2 string
	fmt.Scan(&prefix, &name1, &name2)

	for _, line := range Greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func Greeting(prefix string, name ...string) []string {
	r := []string{}

	for _, v := range name {
		r = append(r, prefix+" "+v)
	}

	return r

}

/*
Extracting digits from a number
Below is a Go program that takes as input two lines: the first line contains a number, and the second line contains three space-separated digits (digit1, digit2, digit3).

Your task is to write the additional required code to create two functions:

convertNumberToSlice() — To convert the input number into a slice of its digits.
suitableDigits() — That takes a slice of digits from a number and a variadic slice of digits to check. This function should return a slice of the digits included in the number.
Finally, your program should print the result of the function suitableDigits() to the console.

Sample Input 1:
1234
1 5 6
Sample Output 1:
[1]
*/

// Write the code to convert the input `number` into a slice of its digits below:
func convertNumberToSlice(number int64) []int8 {
	var digits []int8

	for number > 0 {
		digit := int8(number % 10)
		number /= 10 // Divide by 10 to remove the last digit from the `number`
		digits = append(digits, digit)
	}

	return digits // Return the slice of digits
}

// Write the code to check which of the provided digits are included in the `number` below:
func suitableDigits(numberAsDigits []int8, digitsToCheck ...int8) []int8 {
	var included []int8

	for _, digit := range numberAsDigits {
		for i, digitToCheck := range digitsToCheck {
			if digit == digitToCheck {
				included = append(included, digit)
				digitsToCheck[i] = -1 // Set the found digit to -1 to prevent it from being included again
				break
			}
		}
	}

	return included // Return the slice of included digits
}

// DO NOT delete or modify the contents of the main() function!
func main2() {
	var number int64
	var digit1, digit2, digit3 int8
	fmt.Scan(&number, &digit1, &digit2, &digit3)
	// number = 1234
	// digit1 = 1
	// digit2 = 5
	// digit3 = 6

	numberAsDigits := convertNumberToSlice(number)
	suitable := suitableDigits(numberAsDigits, digit1, digit2, digit3)

	fmt.Println(suitable)
}

// calculateHalf prints the result divided by 2
func calculateHalf(r int) {
	fmt.Println("The result is", r/2)
}

// halfSum sums a and b and then calls calculateHalf to print the result divided by 2
func halfSum(a int, b int) {
	calculateHalf(a + b)
}

// halfDifference subtracts b from a then calls calculateHalf to print the result divided by 2
func halfDifference(a int, b int) {
	calculateHalf(a - b)
}

// DO NOT delete or modify the code within the main function!
func main3() {
	var a, b int
	fmt.Scanln(&a, &b)
	halfSum(a, b)
	halfDifference(a, b)
}

// nolint: gomnd // <-- DO NOT delete this comment!
func fahrenheitToCelsius(temperature float64) {
	celsius := (temperature - 32)	
	celsius /= 1.8
	fmt.Printf("%.2f C\n", celsius)
}

// nolint: gomnd // <-- DO NOT delete this comment!
func celsiusToFahrenheit(temperature float64) {
	fahrenheit := (temperature * 1.8) + 32
	fmt.Printf("%.2f F\n", fahrenheit)
}

// Create a function convertTemperature() that takes `temperature` and `unit` as arguments
// And then calls the correct function to convert the temperature based on the `unit`
func convertTemperature(temperature float64, unit string) {
	if unit == "F" {
		// call the function that converts Fahrenheit to Celsius here!
		fahrenheitToCelsius(temperature)
	}

	if unit == "C" {
		celsiusToFahrenheit(temperature)
	}
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var temperature float64
	var unit string
	fmt.Scanln(&temperature, &unit)

	convertTemperature(temperature, unit)
}
