package main

import (
	"errors"
	"fmt"
	"os"
)

// The function openFile returns a custom error message if opening the file fails
// unsig %w to wrap original os.ReadFile error message
func openFile(filename string) error {
	if _, err := os.ReadFile(filename); err != nil {
		return fmt.Errorf("error opening %s -> %w", filename, err)
	}

	return nil
}

func main() {
	err := openFile("new_file.txt")

	if err != nil {

		fmt.Println(err)
		fmt.Printf("error running main.go: %s\n", err) // Print the wrapped error message

		unwrappedErr := errors.Unwrap(err)                // This line unwraps the error
		fmt.Printf("unwrapped error: %s\n", unwrappedErr) // Print the original error message
	}

	//exemple:
	fmt.Print("\n\n\n")
	err1 := errors.New("error: am I the original error?")
	err1 = fmt.Errorf("error: no! I'm the original error! %w", err1)

	fmt.Println(err1)
	// call the errors.Unwrap function here and pass err as the argument
	unwrappedErr := errors.Unwrap(err1)

	// print the unwrapped error here
	fmt.Print(unwrappedErr)
}
