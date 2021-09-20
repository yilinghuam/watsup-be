package models

import (
	"database/sql"
	"fmt"

	"watsup.com/database"
)

func CheckUserExist(email string) (string, error) {
	db := database.GetInstance()
	// check if user exist
	var user string
	err := db.QueryRow(fmt.Sprintf("SELECT user_id FROM users WHERE email='%s'", email)).Scan(&user)
	// close if there is an error
	if err == sql.ErrNoRows {
		// create user here
		message, Adderr := AddUser(email)
		if Adderr != nil {
			return message, Adderr
		}
		return "", nil
	}
	if err != nil {
		fmt.Println("query issue")
		fmt.Println(err)
		return "", err
	}

	return user, nil
}

func AddUser(email string) (string, error) {
	db := database.GetInstance()
	// check if user exist
	res, err := db.Exec(fmt.Sprintf("INSERT INTO users (user_id,email) VALUES('null','%s')", email))
	// close if there is an error
	if err != nil {
		return "", err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return "", err
	}
	fmt.Printf("The last inserted row id: %d", lastId)

	return "created", nil
}
