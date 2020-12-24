package main

//Data is a struct that holds the time and the value
type Data struct {
	time int
	num  int
}

//Person is a  structure for each animal
type Person struct {
	id    string
	value []Data
}

//Settings holds what the run will do
type Settings struct {
	graph  bool
	folder bool
}
