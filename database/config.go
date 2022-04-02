package database

import "fmt"

type Config struct {
	ServerName string
	User       string
	Password   string
	DBName     string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.ServerName, config.DBName)
	// var connectionString string = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_DBNAME"))

	return connectionString
}
