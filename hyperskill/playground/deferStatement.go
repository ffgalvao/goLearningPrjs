package main

/*
Useful defer statement
Below is a Go program that takes as an input three int values; these values are further read in the hidden function hiddenValuesHandler() and returned as a result.

Your task is to use the defer statement with an anonymous function to check the three resulting values following these rules:

The value of the variable num1 must be in the range from 0 to 100 ([0, 100]);
The value of the variable num2 must be a positive number (num2 > 0);
The value of the variable num3 must be an even number.
Finally, depending on the result of checking each variable, your program should print OK or ERROR.

Sample Input 1:
1 2 3

Sample Output 1:
num1: OK
num2: OK
num3: ERROR
*/


import "fmt"

func main1() {
    var num1, num2, num3 int32

    defer func(){
        if(num1>= 0 && num1 <= 100){
            fmt.Println("num1: OK")
        } else {
            fmt.Println("num1: ERROR")
        }
        
        if(num2> 0){
            fmt.Println("num2: OK")
        } else {
            fmt.Println("num2: ERROR")
        }
    
        if(num3%2 == 0){
            fmt.Println("num3: OK")
        } else {
            fmt.Println("num3: ERROR")
        }
    }()

    // Do not remove the line below!
    // num1, num2, num3 = hiddenValuesHandler()
}

/* 
Using defer to print a reversed string! https://hyperskill.org/learn/step/17699
SpongeBob has been learning about the defer statement! He's created a Go program that takes a name (of the string type) as input and prints it in reverse. Unfortunately, he forgot to add one necessary defer statement and delete some unnecessary ones!

Please help SpongeBob fix his code so his program can work properly. The expected output of his program is the following:

Your name is: SpongeBob
Reversed name: boBegnopS

Sample Input 1:
SpongeBob
Sample Output 1:
Your name is: SpongeBob
Reversed name: boBegnopS 
*/

func main() {
    // delete or add new defer statements below!
    var name string
    fmt.Scanln(&name)

    fmt.Printf("Your name is: %s\n", name)

    for _, char := range name {
        defer fmt.Printf("%c", char)
    }

	defer fmt.Printf("Reversed name: ")

	append()
}