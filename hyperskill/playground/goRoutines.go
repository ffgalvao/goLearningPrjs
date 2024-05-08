package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello, friend!")
}

func main() {
    fmt.Println("Main is started")

    go sayHello()

    fmt.Println("Main is running")
    time.Sleep(time.Millisecond)
    fmt.Println("Main is finished")
}

// Output:
// Main is started
// Main is running
// Hello, friend!
// Main is finished