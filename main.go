package main

import (
	"log"
	"os"
	"phonebook/database"
	"phonebook/entity"

	"github.com/joho/godotenv"
)

func allUsers() {
	var phonebook []entity.Phonebook
	database.Connector.Find(&phonebook)
	log.Printf("%d rows found.", database.Connector.Find(&phonebook).RowsAffected)
	rows, err := database.Connector.Find(&phonebook).Rows()
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		var phonebook entity.Phonebook
		err = database.Connector.ScanRows(rows, &phonebook)
	}
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", phonebook)

	if rows.Err() != nil {
		log.Fatalln(rows.Err())
	}
}
func main() {

	initDB()
	allUsers()

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
