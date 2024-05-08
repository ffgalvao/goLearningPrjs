package main

import "fmt"

func main() {
	// exem01()
	// exem02()
	// exem03()
	// exem04()
	exem05()
}

func exem05() {
	/*
		Advanced examples
		You will not see questions about the following examples in the practice section.
		Both of them are about having a pointer that points to another pointer that points
		to some "solid" type. Sounds complicated, and you'll rarely see such uses in
		practical application, but you can consider these code snippets and play with
		them for a bit. It'll definitely help you with the understanding of pointers.
	*/

	var p **string
	p = new(*string)
	*p = new(string)
	fmt.Println(**p)

	**p = "is this even possible?"
	fmt.Println(**p)

	var s = "yes, it is possible"
	var p1 = &s
	var p2 = &p1

	fmt.Println(**p2)
	fmt.Println(*p2 == p1) // Will print true
	fmt.Println(**p2 == s) // Will print true
}

func exem04() {

	//	Below are proper examples of getting and assigning pointers:

	/*
	   It is important to have the new function in the second example.
	   Otherwise, we would have a nil pointer dereference error at *p = s.
	*/

	var p1 *string
	var s1 = "my string"

	p1 = &s1

	fmt.Println(*p1)

	var p2 = new(string)
	var s2 = "my string"

	*p2 = s2

	fmt.Println(*p2)
}

func exem03() {

	/**
	Remember, the p variable is nothing more than just a fancy integer representing
	the memory address. We cannot assign a string to it directly.
	By putting an asterisk before the variable name, we tell Go that we want to
	work not with the p variable but with the variable at the address held by the
	p pointer. In this case, you cannot put an asterisk anywhere you want,
	only before a valid pointer.
	*/

	var p = new(string)
	*p = "my string"

	fmt.Println(p)
	fmt.Println(*p)
}

func exem02() {
	var p *string
	fmt.Println(*p) // invalid memory address or nil pointer dereference
}

func exem01() {
	var p = new(string)
	fmt.Println(p) // Will print some memory address, for eg. 0xc000040240
	// This address can be different in your case

	fmt.Printf("%#v", *p) // Will print the actual value: "" an empty string
	fmt.Println("")
}

func context() {
	fmt.Println("It's a life")

	/**
	when we just declare a pointer, it doesn't point anywhere.
	To create a pointer of some type and actually allocate memory
	for its default value, we should use a special built-in function new.
	Let's have a look at how we can use this function:
	*/

	var p *string // Declaring p as a pointer to a string

	p = new(string) // Allocating the memory for the default string value
	// and assigning its address to the p pointer

	// This new function returns a pointer to a type, specified as the function argument.
	// We can also merge these two lines into one:

	var p1 = new(string)

	fmt.Println(p, p1)
}
