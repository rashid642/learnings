package routes

import (
	"github.com/gorilla/mux"
	"github.com/rashid642/StudentManagement/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/addPeople", handlers.AddPeople).Methods("POST")
}