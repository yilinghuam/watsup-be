package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"watsup.com/envload"
)

var instance *sql.DB

func GetInstance() *sql.DB {
	return instance
}

func init() {
	var err error

	hostConfig := envload.GoDotEnvVariable("HOST")
	databaseConfig := envload.GoDotEnvVariable("DATABASE")
	UserConfig := envload.GoDotEnvVariable("DB_USER")
	PasswordConfig := envload.GoDotEnvVariable("PASSWORD")
	PortConfig := envload.GoDotEnvVariable("PORT")

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", UserConfig, PasswordConfig, hostConfig, PortConfig, databaseConfig)
	fmt.Println(connectionString)
	// connect to sql
	instance, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("once")
	// to signify connected
	pingErr := instance.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}
