package main

/* 
Greedy editor
Imagine you're a text writer with a unique payment system: you earn 2 cents for each letter and 1 cent for each space you write. To calculate your earnings efficiently, you decide to create a Go program to count the number of letters (symbols excluding spaces) and spaces in a given text.

Your task is to complete the symbolsCounter function, which should return the count of symbols (excluding spaces) in the provided text.

Tip: You can use a function from the strings package to replace all the blank spaces " " with the empty character "", and then count the remaining symbols.

The above method will only work with ASCII text. In it, all characters have a size of 1 byte. In other cases, you can use the utf8 package. Finally, note that the input text only contains ASCII characters for this task!

Sample Input 1:
One day, the noble Jacques, who lived in Vesteros, decided to move to another continent. But Vesteros didn't want to let him go.

Sample Output 1:
Symbols in text: 106
Spaces in text: 22 
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Which function from the `strings` package is used to replace all appearances
// of a word or character in a string with another word or character?
func symbolsCounter(text string) int {
	text = strings.ReplaceAll(text, " ", "")
	return len(text)
}

// DO NOT delete or modify the contents of the main function!
func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	allLength := len(scan.Text())
	symbols := symbolsCounter(scan.Text())
	fmt.Printf("Symbols in text: %d\nSpaces in text: %d", symbols, allLength-symbols)
}