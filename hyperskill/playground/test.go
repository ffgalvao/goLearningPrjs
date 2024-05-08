package main

import (
	"fmt"
	"math"
	. "playground/common/iif"
)

func main() {

	// var a, b int

	// fmt.Println("Enter the numbers:");
	// fmt.Scan(&a)
	// fmt.Scan(&b)

	// fmt.Println(a + b)
	// fmt.Printf("%d %[1]d %d %[1]d %[4]d\n", 1, 2, 3, 4)
	// fmt.Printf("%s %[3]s %s %[1]s %[4]s\n", "a", "b", "c", "d")

	// fmt.Println(Iif(true).Int(2, 3))

	// fmt.Printf("%[1]T %[1]v\n", Iif(true).Int(1, 0))
	// fmt.Printf("%[1]T %[1]v\n", Iif(false).Int(1, 0))
	// fmt.Printf("%[1]T %[1]v\n", Iif(true).Str("1", "0"))
	// fmt.Printf("%[1]T %[1]v\n", Iif(false).Str("1", "0"))
	// fmt.Printf("%[1]T %[1]v\n", Iif(7 > 3).Then("7", 3))

	// b := Iif(7 > 3).Then("5", 3)
	// fmt.Println(b)

	// c := Iif("a" > "c").Str("7", "3")
	// fmt.Println(c)

	// d := Iif2{Cond: 3 > 4, ITrue: 4, IFalse: 775}.Then()
	// fmt.Println(d)

	// d = int(4 + 6)
	// fmt.Println(d)

	fmt.Println("Int")
	fmt.Printf("%[1]T %[1]v\n", IIF(true, 1, 0))
	fmt.Printf("%[1]T %[1]v\n", IIF(false, 1, 0))

	fmt.Println("Str")
	fmt.Printf("%[1]T %[1]v\n", IIF(true, "true", "false"))
	fmt.Printf("%[1]T %[1]v\n", IIF(false, "true", "false"))

	fmt.Println("Float")
	fmt.Printf("%[1]T %[1]v\n", IIF(true, 1.1, 0.0))
	fmt.Printf("%[1]T %[1]v\n", IIF(false, 1.1, 0.0))

	fmt.Println("Struct")
	type strTest struct {
		a string
		b int
	}
	sT := strTest{a: "true", b: 1}
	sF := strTest{a: "False", b: 0}
	fmt.Printf("%[1]T %[1]v\n", IIF(true, sT, sF))
	fmt.Printf("%[1]T %[1]v\n", IIF(false, sT, sF))

	// fmt.Println("Mix")
	// Doenst work
	// fmt.Printf("%[1]T %[1]v\n", IIF(true, "1.1", 0.0))
	// fmt.Printf("%[1]T %[1]v\n", IIF(false, "1.1", 0.0))

}

// Do not change the type declarations below!
type (
	Square struct {
		Side float64
	}

	Circle struct {
		Radius float64
	}

	Shape interface {
		Area() float64
	}
)

// DO NOT change these lines -- they create the proper output string:
func (s Square) String() string { return fmt.Sprintf("Square area: %.2f", s.Area()) }
func (c Circle) String() string { return fmt.Sprintf("Circle area: %.2f", c.Area()) }

// Implement the interface methods for the 'Square' and 'Circle' structs below:
func (s Square) Area() float64 { return s.Side * s.Side }
func (c Circle) Area() float64 { return math.Pi * (c.Radius * c.Radius) }

func main4() {
	// Do NOT change the code below! This reads the input and outputs as required.
	var length float64
	fmt.Scanln(&length)

	for _, shape := range []Shape{Square{length}, Circle{length / 2}} {
		fmt.Println(shape)
	}
}

type Text string

// Create the `TabFormat()` method below:
func (t Text) TabFormat() string {
	return "\t" + string(t)
}

// Create the `DoubleQuoteFormat()` method below:
func (t Text) DoubleQuoteFormat() string {
	return `"` + string(t) + `"`
}

// Create the `SingleQuoteFormat()` method below:
func (t Text) SingleQuoteFormat() string {
	return "'" + string(t) + "'"
}
