package main

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
