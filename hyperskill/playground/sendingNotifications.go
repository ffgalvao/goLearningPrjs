package main

/*
Sending Notifications https://hyperskill.org/learn/step/18733
Suppose you want to create a Notifier interface; it should implement the SendNotification() method that sends a string with a message to a certain email.

We've already created the Notification struct for you. Your task is to create the SendNotification() method and the Notifier interface.

To make the Notifier interface work as intended, the program should take as an input an email, to which it will send the notification, and also the msg that the notification will contain. Finally, it should print the notification in the following format:

Sending notification to username@mail.com with message msg
This problem uses the bufio and os packages to read a whitespace-separated string from the input. To solve this task, you don't need to know how the bufio/os packages work; your only objective is to create the SendNotification() method and the Notifier interface.

Sample Input 1:
impostor@amongus.com
I don't know man, you're looking kind of sus! ඞ
Sample Output 1:
Notification sent to: impostor@amongus.com with message: I don't know man, you're looking kind of sus! ඞ 
*/

import (
	"bufio"
	"fmt"
	"os"
)

type Notification struct {
	Email   string
	Message string
}

// Finish creating the 'SendNotification()' method of the 'string' type below:
func (n Notification) SendNotification() string {
	return fmt.Sprintf("Notification sent to: %s with message: %s", n.Email, n.Message)
}

// Create the 'Notifier' interface that implements the 'SendNotification()' method below.
type Notifier interface {
    SendNotification() string
}

// Do not change the code within the main function!
// The purpose of this task is to create the 'Notifier' interface and 'SendNotification' method above!
func main() {
	var n Notifier

	var email string
	fmt.Scanln(&email)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()

	n = Notification{Email: email, Message: msg}
	fmt.Println(n.SendNotification())
}
