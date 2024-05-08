package main

import "fmt"

func main() {
	var xint uint = 33
	var xFloat32 float32 = 33.3
	var xRune rune = 64
	var xString = "test"

	fint := float32(xint)
	rint := rune(xint)
	sint := string(xint)

	fmt.Println(xint, fint, rint, sint)

	ifloat := int(xFloat32)
	rfloat := rune(xFloat32)
	// sfloat := string(xFloat32)

	fmt.Println(xFloat32, ifloat, rfloat)

	irune := int(xRune)
	frune := float32(xRune)
	srune := string(xRune)

	fmt.Println(xRune, irune, frune, srune)

	// iStr := int(xString)
	// fStr := float32(xString)
	rStr := []rune(xString)

	fmt.Println(xString, rStr)

}
