package main
/* 
Creating a slot machine in Go
Edwin wants to create a three-slot machine in Go. To do this, he has created the SlotMachine struct and has written some code within the main() function so users of the slot machine can enter the betAmount.

Now Edwin needs to create two methods to make the slot machine work:

Play() — It should use a for loop to iterate over the s.Slots array and assign a random int number to each element of the slice using the rand.Intn() function.

CheckWin() — It should check if all three elements of the s.Slots array are equal and return the message: "Jackpot!". Otherwise it should return the message "You lost!".

This problem uses the math/rand package to generate random numbers for the slot machine output. However, don't be scared! You don't need to know how the math/rand package works to solve this problem. Your only objective is to write the additional required code to finish the implementation of the Play() and CheckWin() methods.

Sample Input 1:
20
Sample Output 1:
[8 3 5]
You lost!

Sample Input 2:
710
Sample Output 2:
[7 7 7]
Jackpot!

Sample Input 3:
287
Sample Output 3:
[5 5 5]
Jackpot! 
*/

import (
	"fmt"
	"math/rand"
)

type SlotMachine struct {
	Slots [3]int
}

const nineUnits = 9

// Create the Play() and CheckWin() methods below:
func (s *SlotMachine) Play() {
	for i := range s.Slots {
		s.Slots[i] = rand.Intn(nineUnits)
	}
}

func (s *SlotMachine) CheckWin() string {
	if s.Slots[0] == s.Slots[2] && s.Slots[1] == s.Slots[2] {
		return "Jackpot!"
	}
	return "You lost!"
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var money int64
	fmt.Scanln(&money)
	r := rand.New(rand.NewSource(money))
	r.Seed(money)
	// rand.Seed(money)

	var sm SlotMachine
	sm.Play()
	fmt.Println(sm.Slots)
	fmt.Println(sm.CheckWin())
}