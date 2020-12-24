package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//file is temporary

var people []Person

//GetPeopleEndpoint handles endpoint
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)

}

//Router creates endpoints
func Router() {

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")

}
