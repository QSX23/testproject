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

func main() {
	file := ImportCSVFile("placeholder.csv")
	data := ReadData(file)
	people := ParseLines(data)

	//iteratation testing
	/*for _, v := range people {
		for _, x := range v.value {
			fmt.Println(x)
		}
		fmt.Println(v)
	}*/

	GraphData(people)

}
