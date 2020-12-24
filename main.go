package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

//conditional for the graph

var settings Settings

func main() {

	//reads specified file from directory
	file := ImportCSVFile("placeholder.csv")
	data := ReadData(file)
	people := ParseLines(data)
	settings.graph = true
	settings.folder = true

	//reads all files from data folder
	files := ReadFiles()

	for _, file := range files {
		fmt.Println(file)
	}

	//iteratation testing
	/*for _, v := range people {
		for _, x := range v.value {
			fmt.Println(x)
		}
		fmt.Println(v)
	}*/

	if settings.folder == true {
		fullData := OpenAll(files)
		//fmt.Println(fullData)

		if settings.graph == true {
			for k, v := range fullData {
				name := strings.TrimSuffix(k, filepath.Ext(k))
				fmt.Println(name)
				GraphData(v, name)

			}
		}
	} else if settings.graph == true {
		GraphData(people, file.Name())
	}

}
