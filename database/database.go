package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var instance *sql.DB

func GetInstance() *sql.DB {
	return instance
}

func init() {
	// Initalize the sql.DB connection pool and assign it to the models.DB
	var err error
	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   "lingy@watsup",  //os.Getenv("DBUSER")
	// 	Passwd: "Chipmunk@001", //os.Getenv("DBPASS")
	// 	Net:    "tcp",
	// 	Addr:   "watsup.mysql.database.azure.com",
	// 	DBName: "",
	// }
	const (
		host     = "watsup.mysql.database.azure.com"
		database = "watsup"
		user     = "lingy@watsup"
		password = "Chipmunk@001"
	)
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

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
