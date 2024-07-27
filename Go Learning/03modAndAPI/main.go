package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Id        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name`
	Address   *Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []Person

func GetPeople(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{
		Id:        1,
		FirstName: "Rashid",
		LastName:  "Aziz",
		Address: &Address{
			City:  "Mumbai",
			State: "Maharastra",
		},
	})
	people = append(people, Person{
		Id:        2,
		FirstName: "Yasin",
		LastName:  "Khan",
		Address: &Address{
			City:  "Mumbai",
			State: "Maharastra",
		},
	})

	router.HandleFunc("/people", GetPeople).Methods("get")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// go mod init github.com/rashid642/modLearning
