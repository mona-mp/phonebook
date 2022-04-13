package main

import (
	"log"
	"net/http"
	"os"

	"phonebook/controllers"
	"phonebook/database"
	"phonebook/entity"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 18080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":18080", router))

}

func initaliseHandlers(router *mux.Router) {

	router.HandleFunc("/phonebook/users/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/phonebook/users", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/phonebook/users/{id}", controllers.DeletPersonByID).Methods("DELETE")
	router.HandleFunc("/phonebook/users", controllers.GetAllPerson).Methods("GET")

}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config :=
		database.Config{
			ServerName: os.Getenv("DB_SERVER"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_DBNAME"),
		}
	connetionString := database.GetConnectionString(config)
	err = database.Connect(connetionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Phonebook{})
}
