package main

/*
Simple Password Generator
Let's create a simple Password Generator that consists of 4 symbols: an uppercase letter, a lowercase letter, a number, and a special character.

The input is the Seed number, and the output is the string in the form of UpperLowerNumberCharacter, for example "Bd3$".

Sample Input 1:
1
Sample Output 1:
Gv7#
*/

import (
	"fmt"
	"math/rand"
)

func main1() {
	upperCase := "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	// put your code here

	pwd := make([]byte, 4)

	pwd[0] = upperCase[rand.Intn(len(upperCase))]
	pwd[1] = lowerCase[rand.Intn(len(lowerCase))]
	pwd[2] = number[rand.Intn(len(number))]
	pwd[3] = specialSymbol[rand.Intn(len(specialSymbol))]
	rand.Float32()

	fmt.Println(string(pwd))

}

/*
Coin Toss https://hyperskill.org/repeat/step/28504
Let's create a project that imitates a coin toss.

Using a random number generator create a simple project where the output is TAILS if the random number is between [0, 0.5) and HEADS if it is between [0.5, 1.0).

The input is the Seed number, and the output is either TAILS or HEADS according to the abovementioned conditions.

Sample Input 1:
10
Sample Output 1:
HEADS
*/

func main2() {
	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	// put your code here
	if rand.Float32() <= 0.5 {
		fmt.Println("TAILS")
	} else {
		fmt.Println("HEADS")
	}

}

/*
Random email addresses https://hyperskill.org/repeat/step/28506

Let's suppose you work for a company, and your boss requests you to generate 3 email addresses with
five random lowercase letters following the domain name @fantasy.com.

Note: To generate a random string consisting of 5 random lowercase letters you can use ASCII numbers

The input is the Seed number, and the output is three email addresses, for example ghjkr@fantasy.com .

Sample Input 1:
1

Sample Output 1:
xvlbz@fantasy.com
gbaic@fantasy.com
mrajw@fantasy.com

*/

func main() {
	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	// put your code here

	lowerCase := "abcdefghijklmnopqrstuvwxyz"

	emails := make([]string, 3)

	for i := 0; i < cap(emails); i++ {
		email := make([]byte, 5)

		for j := 0; j < cap(email); j++ {
			email[j] = lowerCase[rand.Intn(len(lowerCase))]
		}

		emails[i] = string(email) + `@fantasy.com`
	}

	for _, val := range emails {
		fmt.Println(val)
	}

}
