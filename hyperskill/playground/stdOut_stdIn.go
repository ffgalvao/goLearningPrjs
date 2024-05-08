package main

import "fmt"

func main() {

	// exem01();
	// exem02();
	// exem03()
	exem04()

}

func exem04() {
	var name string
	var age int

	fmt.Print("Enter your name: ") // Writing a request message to the stdout
	fmt.Scan(&name)                // Reading from the stdin into the name variable
	fmt.Println("")                // Going to the next line by writing /n to the stdout

	fmt.Print("Enter your age: ") // Writing a request message to the stdout
	fmt.Scan(&age)                // Reading from the stdin into the age variable
	fmt.Println("")               // Going to the next line by writing /n to the stdout

	fmt.Println(name, age) // Writing to the stdout the values of name and
	// age variables that you have entered
}

func exem03() {
	var name string
	var age string

	fmt.Println("Enter you name and age:")
	fmt.Scan(&name, &age) // Reading from the stdin into the name and age variables

	fmt.Println(name) // Writing to the stdout the value of the name variable
	fmt.Println(age)  // Writing to the stdout the value of the age variable
}

func exem02() {
	var foo int    // foo is 0
	var str string // string is ""

	fmt.Println("If you enter a string character, the scan function")
	fmt.Scan(&foo) // If you enter a string character, the scan function
	// will leave the variable foo unchanged

	fmt.Println("If you enter an integer, it will be taken as a string")
	fmt.Scan(&str) // If you enter an integer, it will be taken as a string
	// and assigned to the variable str

	fmt.Println(foo, str)
}

func exem01() {
	var name string

	fmt.Println("My first Print Out")
	fmt.Println("Make me happy, tell me your name?")

	fmt.Scan(&name)

	fmt.Println("Thanks " + name)
}
