package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

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
	m := make([]Person, len(lines[0]))

	for int := 1; int < 7; int++ {
		id := lines[0][int]
		m[int] = Person{
			id: id,
		}

		for _, line := range lines {
			t, err := strconv.Atoi(line[int])
			if err != nil {
				t = 0
			}
			d := Data{
				time: line[0],
				num:  t,
			}
			m[int].value = append(m[int].value, d)
		}
	}

	return m
}

func main() {
	file := ImportCSVFile("placeholder.csv")
	data := ReadData(file)
	people := ParseLines(data)

	for _, v := range people {
		fmt.Println(v)
	}

}
