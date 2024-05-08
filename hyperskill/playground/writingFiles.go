package main

/*
Keanu is writing his most famous quotes!
Keanu Reeves wants to write three of his most famous quotes contained in the slice of
strings quotes to the keanu_quotes.txt file.

Can you help Keanu implement the additional code required to write strings line by line to
the file?

Take notice that Keanu has already written the code required to open keanu_quotes.txt in
read-and-write mode and assigned it to the file variable.
*/

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main2() {
	// DO NOT DELETE! This opens the 'keanu_quotes.txt' file in read-write mode
	filepath, _ := filepath.Abs("files/keanu_quotes.txt")
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	quotes := []string{"You're breathtaking!",
		"Wake up Samurai! We have a city to burn.",
		"Lose? I don't lose; I win! I'm a lawyer, that's my job, that's what I do!",
	}

	for _, line := range quotes { // iterate over each line of the 'quotes' slice here
		_, err := fmt.Fprintln(file, line) // // write each line of the 'quotes' slice of strings to 'file' here
		if err != nil {
			log.Fatal(err) // exit if we have an unexpected error
		}
	}
}

/*
Writing songs to a CSV file https://hyperskill.org/learn/step/23520
Another common file format apart from .txt files are the .csv files — CSV stands for Comma Separated Values.

A CSV file stores values that are separated by a specific delimiter, commonly a comma (,). However,
there are other common delimiters like the semicolon (;), a tab (\t), blank spaces (" ")
and even the pipe character (|).
Below you can see the contents of a sample CSV file:

Title,Artist,Album,Duration
"Anti-Hero","Taylor Swift","Midnights",3:21
"Midnight Rain","Taylor Swift","Midnights",2:54
Now that you've been acquainted with CSV files, your task is to create a Go program that writes
the comma-separated values contained within the data variable into the songs.csv file.

To do this, you can use the os.WriteFile() function to write the data directly into songs.csv,
or you can use the os.Create() function to create songs.csv first and then use fmt.Fprintln()
to write the data into the file.
*/

func main() {
	// DO NOT modify the contents of the `data` variable!
	data := `Title,Artist,Album,Duration
"Fix You","Coldplay","X&Y",4:56
"Clocks","Coldplay","A Rush of Blood to the Head",5:08
"Yellow","Coldplay","Parachutes",4:27
"Summertime Sadness","Lana Del Rey","Born to Die",4:25
"Young and Beautiful","Lana Del Rey","Born to Die",3:56
"Pumped Up Kicks","Foster the People","Torches",3:59
`

	// Write the data to the "songs.csv" file below
	// You can use the os.WriteFile() to write the data directly to "songs.csv"
	// Or you can use the os.Create() to create "songs.csv" and then use fmt.Fprintln() to write the data.

	filepath, _ := filepath.Abs("/files/songs.csv")
	if err := os.WriteFile(filepath, []byte(data), 0644); err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

}
