package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rashid642/StudentManagement/handlers"
	"github.com/rashid642/StudentManagement/models"
	"github.com/rashid642/StudentManagement/routes"
	"github.com/rashid642/StudentManagement/utils"
)

func main() {
	handlers.People = append(handlers.People, models.Person{
		Id:        1,
		FirstName: "Rashid",
		LastName:  "Aziz",
		Address: &models.Address{
			City:  "Mumbai",
			State: "Maharastra",
		},
	})
	handlers.People = append(handlers.People, models.Person{
		Id:        2,
		FirstName: "Yasin",
		LastName:  "Khan",
		Address: &models.Address{
			City:  "Mumbai",
			State: "Maharastra",
		},
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env file can not be opened")
	}
	MONGO_URI := os.Getenv("MONGODB_URI")
	utils.InitMongoDB(MONGO_URI)

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Println("Server at https://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}