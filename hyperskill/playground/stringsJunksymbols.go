/*
Junk symbols

Text that you copy from the internet often contains junk symbols. Examples are double whitespace, the symbol "minus" instead of a dash, single quotes instead of double ones, or single punctuation marks (dot, comma, exclamation mark, or question mark).

You need to write a small program that points out these junk symbols. If any of the cases are found in the input string, print the corresponding message:

"Double whitespace"
"Minus symbol" - (dash is —)
"Single quotes" '
"Single mark" .,?!
For the minus-dash problem, let's make a deal that the minus symbol means a negative number, and it's written without whitespace: -1, -0.5, -433. On the other hand, the dash symbol has whitespace on both sides: — .

For the single-mark problem, remember that punctuation marks are placed right next to another symbol:

single-mark problem, — correct
single-mark problem , — not correct
And let's think that a string doesn't contain words like it's, where a single quote is a part of the word.

Sample Input 1:

Hello World!
Sample Output 1:
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// var source string = "-Hello World!"
	var source string = " - single-mark problem ,"
	// scanner := "Hello World!"
	//  scanner := bufio.NewScanner(os.Stdin)
	//  if scanner.Scan() {
	// 	 source = scanner.Text()
	//  }

	junk := map[string]string{
		"  ": "Double whitespace",
		" '": "Single quotes",
		" .": "Single mark",
		" ,": "Single mark",
		" ?": "Single mark",
		" !": "Single mark",
	}

	for key, val := range junk {
		if strings.Contains(source, key) {
			fmt.Println(val)
		}

	}
	if strings.Contains(source, " - ") {
		index := strings.Index(source, "-")
		index++
		fmt.Println(string(source[index]))
		if !unicode.IsDigit(rune(source[index])) {
			fmt.Println("Minus symbol")
		}
	}

	//  if strings.Contains(?, ?) {
	// 	 fmt.Println("Double whitespace")
	//  }

	// add the same checks for other cases here or below this comment!
}

func main2() {
	var source string

	fmt.Scan(&source)

	zeros := strings.Count(source, "0")
	ones := strings.Count(source, "1")

	switch {
	case zeros == ones:
		fmt.Println("Counts are equal")
	case zeros > ones:
		fmt.Println("The count of zeros is bigger")
	case ones > zeros:
		fmt.Println("The count of units is bigger")
	default:
		fmt.Println("Invalid")
	}

}
