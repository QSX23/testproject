package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//ReadFiles reads all the files in the data directory. Future function will pass each to importCSVFile
func ReadFiles() []string {

	//get the directory of the project
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root = filepath.Join(root, "data")

	//set those files as a slice of strings
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println("something went wrong")
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files
}

//OpenAll opens all files within the data directory
func OpenAll(files []string) map[string][]Person {

	p := make(map[string][]Person)

	for _, file := range files {
		path := filepath.Join(directory(), file)
		f, err := os.Open(path)
		if err != nil {
			Exit("This is not a file")
		}
		data := ReadData(f)
		m := ParseLines(data)

		p[file] = m

	}

	return p
}

//ImportCSVFile is a function that imports the named csv file this can be flagged from the terminal
func ImportCSVFile(file string) *os.File {
	filename := flag.String("file", file, "the file to be procesed")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		Exit("This is not a file")
	}
	return f
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

//Exit is a helper to close the program
func Exit(x string) {
	fmt.Println(x)
	os.Exit(1)
}

func directory() string {
	//get the directory of the project
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	s := filepath.Join(root, "data")

	return s
}
