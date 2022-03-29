package main

import (
	// "database/sql"

	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Person struct {
	Id          int
	Name        string
	Phonenumber string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_DBNAME"))
	// var GetConnectionString = func(config Config) string {
	// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	// 	return connectionString
	// }

	fmt.Println(os.Getenv("DB_USER"))
	fmt.Print(connectionString)
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connectionString)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO phonebook (name,phonenumber) values ('milad-4',09352631)")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
