package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var instance *sql.DB

func GetInstance() *sql.DB {
	return instance
}

func init() {
	// Initalize the sql.DB connection pool and assign it to the models.DB
	var err error
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",  //os.Getenv("DBUSER")
		Passwd: "mysql", //os.Getenv("DBPASS")
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "watsup",
	}

	// connect to sql
	instance, err = sql.Open("mysql", cfg.FormatDSN())
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
