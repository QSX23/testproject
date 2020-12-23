package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

//ParseLines takes each line and assigns the values to each animal
func ParseLines(lines [][]string) []Person {
	m := make([]Person, len(lines[0])-1)

	for int := 1; int < 7; int++ {
		id := lines[0][int]
		m[int-1] = Person{
			id: id,
		}
		for _, line := range lines {
			n, err := strconv.Atoi(line[int])
			if err != nil {
				n = 0
			}
			t, err := strconv.Atoi(line[0])

			if err != nil {
				t = 0
			}

			d := Data{
				time: t,
				num:  n,
			}
			m[int-1].value = append(m[int-1].value, d)
		}
	}
	return m
}
