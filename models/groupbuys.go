package models

import (
	"fmt"

	"watsup.com/database"
)

type Groupbuy struct {
	Groupbuy_id      int16
	User_id          string
	Name             string
	Description      string
	Order_date       string
	Closing_date     string
	Delivery_options bool
	Delivery_price   int64
	Status           string
}

type CreateGroupbuyConfig struct {
	User_id          string
	Name             string
	Description      string
	Order_date       string
	Closing_date     string
	Delivery_options bool
	Delivery_price   int64
}

type GroupbuyDashboard struct {
	Name       string `json:"name"`
	Order_date string `json:"order_date"`
	Status     string `json:"status"`
}

func CreateGroupbuy(table CreateGroupbuyConfig) error {
	fmt.Println(table)
	// for checking form info
	fmt.Printf("INSERT INTO watsup.groupbuy(user_id,name,description,order_date,closing_date,delivery_options,delivery_price,status)VALUES(%s,%s,%s,%s,%s,%t,%d)", table.User_id, table.Name, table.Description, table.Order_date, table.Closing_date, table.Delivery_options, table.Delivery_price)
	rows, err := database.GetInstance().Query(fmt.Sprintf("INSERT INTO watsup.groupbuy(user_id,name,description,order_date,closing_date,delivery_options,delivery_price)VALUES('%s','%s','%s','%s','%s',%t,%d)", table.User_id, table.Name, table.Description, table.Order_date, table.Closing_date, table.Delivery_options, table.Delivery_price))
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	fmt.Println(rows)

	return nil
	// item table, groupbuy table
}

func GroupbuyByHost(user string) ([]GroupbuyDashboard, error) {
	fmt.Println(user)
	rows, err := database.GetInstance().Query(fmt.Sprintf("SELECT name,order_date,status FROM groupbuy WHERE user_id='%s'", user))
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return nil, err
	}
	defer rows.Close()

	var lists []GroupbuyDashboard

	//for loop to read through each row of data
	for rows.Next() {
		var list GroupbuyDashboard
		err := rows.Scan(&list.Name, &list.Order_date, &list.Status)
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

func AllOpenGroupbuy(user string) ([]Groupbuy, error) {
	fmt.Println(user)
	rows, err := database.GetInstance().Query(fmt.Sprintf("SELECT * FROM groupbuy WHERE user_id='%s'", user))
	fmt.Println(rows)
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return nil, err
	}
	defer rows.Close()

	var lists []Groupbuy

	//for loop to read through each row of data
	for rows.Next() {
		var list Groupbuy
		err := rows.Scan(&list.Groupbuy_id, &list.User_id, &list.Name, &list.Description, &list.Order_date, &list.Closing_date, &list.Delivery_options, &list.Delivery_price, &list.Status)
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
