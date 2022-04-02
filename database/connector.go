// package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// var db = func() *sql.DB {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_DBNAME"))

// 	db, err := sql.Open(os.Getenv("DB_DRIVER"), connectionString)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	return db
// }
package database

import (
	"log"
	"phonebook/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

var Connector *gorm.DB

func Connect(connectionString string) error {

	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	// Connector.LogMode(true)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!")
	return nil
}

func Migrate(table *entity.Phonebook) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
