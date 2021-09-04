package models

import (
	"fmt"

	"watsup.com/database"
)

// id  INT AUTO_INCREMENT NOT NULL,
//     name VARCHAR(255) NOT NULL,
//     description VARCHAR(255),
//     frequency   INT NOT NULL,
// 	cycle       VARCHAR(255) NOT NULL,
// 	stars       INT NOT NULL,
//     user   VARCHAR(255) NOT NULL,
//     current_frequency INT DEFAULT 0,
//     completed VARCHAR(255) DEFAULT 'not',
//     start_date TIMESTAMP,
//     end_date TIMESTAMP,
//     PRIMARY KEY(id)
type List struct {
	Id          int
	Category    string `json:"cateogory"`
	Main_goal   string `json:"main_goal"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Frequency   int8   `json:"frequency"`
	Cycle       string `json:"cycle"`
	Stars       int8   `json:"stars"`
	User        string `json:"user"`
	State       string `json:"state"`
}

func AllLists() ([]List, error) {
	rows, err := database.GetInstance().Query("SELECT * FROM lists")
	fmt.Println(rows)
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return nil, err
	}
	defer rows.Close()

	var lists []List

	//for loop to read through each row of data
	for rows.Next() {
		var list List
		err := rows.Scan(&list.Id, &list.Category, &list.Main_goal, &list.Name, &list.Description, &list.Frequency, &list.Cycle, &list.Stars, &list.User, &list.State)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return nil, err
		}

		lists = append(lists, list)

	}
	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	return lists, nil
}

func SingleList(listName string) ([]List, error) {
	user := "ling"

	rows, err := database.GetInstance().Query(fmt.Sprintf("SELECT * FROM lists WHERE name='%s' AND user='%s'", listName, user))

	if err != nil {
		fmt.Println("single query error")
		return nil, err
	}

	defer rows.Close()

	var lists []List

	for rows.Next() {
		var list List
		err := rows.Scan(&list.Id, &list.Category, &list.Main_goal, &list.Name, &list.Description, &list.Frequency, &list.Cycle, &list.Stars, &list.User, &list.State)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		lists = append(lists, list)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	return lists, nil
}
