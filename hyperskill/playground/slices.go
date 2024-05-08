/**
Removing all letters except for the last one
Suppose you have a Go program that takes as an input a random word of the string type.

Your task is to use slice expressions along with the built-in len() function to remove all the letters from the word, except for the last one.

For example, if the input word is uncopyrightable, your program should only print e.
**/

package main

import "fmt"

func main() {
	// DO NOT delete or modify the code block below, it reads a random input word
	var word string
	fmt.Scan(&word)

	// Use slice expressions with the len() function 
	// to remove all the letters from 'word' except for the last letter.
	word = word[len(word)-1:]

    fmt.Println(word)


	words := []string{"saltwater", "backstage", "bedrock", "sandcastle", "snowman"}

	w1 := words[0][4:]
	w2 := words[1][4:]
	w3 := words[2][3:]
	w4 := words[3][4:]
	w5 := words[4][3:]

	fmt.Println(w1, w2, w3, w4, w5) // DO NOT delete this line, it prints the substrings!
}