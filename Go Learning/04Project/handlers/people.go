package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rashid642/StudentManagement/models"
	"github.com/rashid642/StudentManagement/utils"
)

var People []models.Person

func GetPeople(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(People)
}

func AddPeople(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside add people API", req.Body)
	var person models.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	_, err := utils.Collection.InsertOne(context.TODO(), person)
	if err != nil {
		http.Error(w, "Error in inserting the data", http.StatusInternalServerError)
		return;
	}

	json.NewEncoder(w).Encode(person);

}