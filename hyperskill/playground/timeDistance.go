package main

/*
Time distance https://hyperskill.org/learn/step/19302
Below is a Go program that takes a year, month and day values as input, creates a date using those values, and uses the now variable to represent the current time.

Your task is to compare the date and now variables using the After() and Sub() functions and print a string according to the following conditions:

"This date will be in the future" — if the input date is any date after the current time (now);

"This date was in the past hundred years" — if the time difference between now and date is less than the current centuryDuration;

"It was a long time ago" — if none of the above conditions are met.

Sample Input 1:
2151 6 16
Sample Output 1:
This date will be in the future

Sample Input 2:
2010 12 4
Sample Output 2:
This date was in the past hundred years

Sample Input 3:
1 1 1
Sample Output 3:
It was a long time ago...


*/

import (
	"fmt"
	"time"
)

const (
	HoursPerDay = 24
	DaysPerYear = 365
	Century     = 100
)

func main() {
	Timezones()
}

func matimeDistance() {
	var year, month, day int
	fmt.Scan(&year, &month, &day)

	// Create the `date` using the input `year`, `month`, and `day`:
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	yearDuration := time.Hour * HoursPerDay * DaysPerYear
	centuryDuration := yearDuration * Century

	now := time.Now()

	// Write the missing code below to compare the `date` and `now` variables:
	switch {
	case date.After(now):
		fmt.Println("This date will be in the future")
	case now.Sub(date) < centuryDuration:
		fmt.Println("This date was in the past hundred years")
	default:
		fmt.Println("It was a long time ago...")
	}
}

func NewYear() {
	now := time.Now()

	nextYear := now.Year() + 1

	firstJanuary := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, time.Local)

	fmt.Println(firstJanuary)
	// The hidden 'check' function checks your answer:
	// check(firstJanuary.Sub(now))
}

const (
	Year  = 1902
	Month = 11
	Day   = 2
)

/*
For a given date, print its value in the most popular timezones:

London (UTC+1.00)
Moscow (UTC+3.00)
New York (UTC-4.00)
Hong Kong (UTC+8.00)
Tokyo (UTC+9.00)
Darwin (UTC+9.30)
Sample Input 1:

Sample Output 1:

London: 1902-11-02 01:00:00 +0000 UTC
Moscow: 1902-11-02 03:00:00 +0000 UTC
New York: 1902-11-01 20:00:00 +0000 UTC
Hong Kong: 1902-11-02 08:00:00 +0000 UTC
Tokyo: 1902-11-02 09:00:00 +0000 UTC
Darwin: 1902-11-02 09:30:00 +0000 UTC
*/

// DO NOT delete the comments on each line below!
func Timezones() {
	
	
	utc := time.Date(Year, Month, Day, 0, 0, 0, 0, time.UTC)
	london := utc.Add(1 * time.Hour)      // nolint: gomnd
	moscow := utc.Add(3 * time.Hour)      // nolint: gomnd
	newYork := utc.Add(-4 * time.Hour)    // nolint: gomnd
	hongKong := utc.Add(8 * time.Hour)    // nolint: gomnd
	tokyo := utc.Add(9 * time.Hour)       // nolint: gomnd
	darwin := tokyo.Add(30 * time.Minute) // nolint: gomnd

	fmt.Printf("London: %v\n", london)
	fmt.Printf("Moscow: %v\n", moscow)
	fmt.Printf("New York: %v\n", newYork)
	fmt.Printf("Hong Kong: %v\n", hongKong)
	fmt.Printf("Tokyo: %v\n", tokyo)
	fmt.Printf("Darwin: %v\n", darwin)
	//Repeat for:
	//Moscow
	//New York
	//Hong Kong
	//Tokyo
	//Darwin
}
