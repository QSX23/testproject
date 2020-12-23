package main

import (
	"flag"
	"fmt"
	"os"
)

//Exit is a helper to close the program
func Exit(x string) {
	fmt.Println(x)
	os.Exit(1)
}

//ImportCSVFile is a function that imports the named csv file this can be flagged from the terminal
func ImportCSVFile(s string) *os.File {
	filename := flag.String("file", s, "the file to be procesed")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		Exit("This is not a file")
	}
	return file
}
