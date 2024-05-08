package main

/*
Server response time conversion https://hyperskill.org/learn/step/32209
Your friend Kenny is currently analyzing server response times. His Go program receives a log that lists the
response times in milliseconds, e.g., "1337ms" as a string, and using slice expressions, extracts the numeric part ("1337")
of the log.

However, Kenny needs the response times in seconds for easier analysis. Your task is to help Kenny write the additional
required code to convert the response time from string to a float64 representation of seconds and then print it out with
two decimal points.

Tip: You can divide the responseTime milliseconds by 1000 to convert them into seconds.

This task uses slice expressions to remove the "ms" suffix from the input log. However, don't be scared!
You don't need to know how slice expressions work to solve this task. Your only objective is to convert the
response time to seconds, represented as a float64 type.

Sample Input 1:
1337ms

Sample Output 1:
Response time: 1.34 seconds

*/


import (
	"fmt"
    "log"
	"strconv"
)

func main() {
    // DO NOT delete or modify the code block below!
	var logEntry string = `1337ms`
	// fmt.Scanln(&logEntry)
	logEntry = logEntry[:len(logEntry)-2]

    // Write the additional code to convert the `responseTime` to float64:
	responseTime, err := strconv.ParseFloat(logEntry, 64)
	if err != nil {
        log.Fatal(err)
	}

    // Convert the response time to seconds below:
    responseTime /= 1000

	// Print the converted response time in seconds with two decimal points.
	fmt.Printf("Response time: %.2f seconds", responseTime)
}
