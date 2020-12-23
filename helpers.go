package main

import (
	"encoding/csv"
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

//ReadData reads the file and returns it as needed
func ReadData(file *os.File) [][]string {
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1
	data, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		Exit("Data could not be read")
	}

	//trim the columns and set the time
	trim := len(data[0])
	v := 0

	for i := range data {
		data[i] = data[i][:trim]
		data[i][0] = strconv.Itoa(v)
		v += 30
	}

	return data
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
