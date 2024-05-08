package main

import (
	"fmt"
	"reflect"
)

var (
	name string
	clip int
)

var (
	name2 = "Kubernetes"
	clip2 = 2
)

func main() {
	fmt.Println(name)
	fmt.Println(clip)

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(clip))

	// Vars 2
	fmt.Println(name2)
	fmt.Println(clip2)

	fmt.Println(reflect.TypeOf(name2))
	fmt.Println(reflect.TypeOf(clip2))

	// Sort Variables
	sayhi := "hello world"
	fmt.Println(sayhi)

	// Full Variables
	var sayhi2 string = "hello world"
	fmt.Println(sayhi2)

}
