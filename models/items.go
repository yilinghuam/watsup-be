package models

import (
	"fmt"
	"strings"

	"watsup.com/database"
)

type ItemConfig struct {
	Id          int8
	User_id     string
	Item        string
	Price       int8
	Quantity    int16
	Groupbuy_id int16
	Order_id    int16
}
type AddItemConfig struct {
	Item        string `json:"Item"`
	Price       int8   `json:"Price"`
	Quantity    int16  `json:"Quantity"`
	Groupbuy_id int16  `json:"Groupbuy_id"`
}
type DashboardSetupConfig struct {
	Item     string `json:"item"`
	Price    int8   `json:"price"`
	Quantity int16  `json:"quantity"`
}

func DeleteItemsByOrderId(order_id int) error {
	_, err := database.GetInstance().Exec("DELETE FROM item where order_id =?", order_id)
	if err != nil {
		fmt.Println("delete items issue")
		return err
	}

	return nil
}
func DeleteItemsByGroupbuy(groupbuy_id int) error {
	_, err := database.GetInstance().Exec("DELETE FROM item where groupbuy_id =?", groupbuy_id)
	if err != nil {
		fmt.Println("delete items issue")
		return err
	}

	return nil
}

func GetUserItemsByUser(user string, groupbuy_id int) ([]ItemConfig, error) {
	db := database.GetInstance()
	// retrieve groupbuy_id
	var items []ItemConfig

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM item WHERE groupbuy_id='%d' and user_id='%s'", user, groupbuy_id))
	// close if there is an error
	if err != nil {
		fmt.Println(err)
		fmt.Println("query issue")
		return items, err
	}
	defer rows.Close()
	fmt.Println(rows)

	//for loop to read through each row of data
	for rows.Next() {
		var item ItemConfig
		err := rows.Scan(&item.Id, &item.User_id, &item.Item, &item.Price, &item.Quantity, &item.Groupbuy_id, &item.Order_id)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(err)
		items = append(items, item)

	}
	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	fmt.Println(err)
	fmt.Println(items)
	return items, nil
}

func GetUserItems(groupbuy_id int) ([]ItemConfig, error) {
	db := database.GetInstance()
	// retrieve groupbuy_id
	var items []ItemConfig

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM item WHERE groupbuy_id='%d' and NOT order_id='0'", groupbuy_id))
	// close if there is an error
	if err != nil {
		fmt.Println(err)
		fmt.Println("query issue")
		return items, err
	}
	defer rows.Close()
	fmt.Println(rows)

	//for loop to read through each row of data
	for rows.Next() {
		var item ItemConfig
		err := rows.Scan(&item.Id, &item.User_id, &item.Item, &item.Price, &item.Quantity, &item.Groupbuy_id, &item.Order_id)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(err)
		items = append(items, item)

	}
	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	fmt.Println(err)
	fmt.Println(items)
	return items, nil
}

func GetHostItems(user string, groupbuy_id int) ([]ItemConfig, error) {
	db := database.GetInstance()
	// retrieve groupbuy_id
	var items []ItemConfig

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM item WHERE user_id='%s' and groupbuy_id='%d' and order_id='0'", user, groupbuy_id))
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return items, err
	}
	defer rows.Close()
	fmt.Println(rows)

	//for loop to read through each row of data
	for rows.Next() {
		var item ItemConfig
		err := rows.Scan(&item.Id, &item.User_id, &item.Item, &item.Price, &item.Quantity, &item.Groupbuy_id, &item.Order_id)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(err)
		items = append(items, item)

	}
	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	fmt.Println(err)
	fmt.Println(items)
	return items, nil
}

func CreateGroupbuyItems(user string, groupbuyName string, data []DashboardSetupConfig) error {
	db := database.GetInstance()
	// retrieve groupbuy_id
	rows, err := db.Query(fmt.Sprintf("SELECT groupbuy_id FROM groupbuy WHERE user_id='%s' and name='%s'", user, groupbuyName))
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return err
	}
	defer rows.Close()

	var groupbuyID int16

	//for loop to read through each row of data
	for rows.Next() {
		err := rows.Scan(&groupbuyID)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return err
		}
	}

	sqlStr := "INSERT INTO item(user_id,item,price,quantity, groupbuy_id) VALUES "
	vals := []interface{}{}
	for _, row := range data {
		sqlStr += "(?,?, ?, ?,?),"
		vals = append(vals, user, row.Item, row.Price, row.Quantity, groupbuyID)
	}
	//trim the last ,
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	//prepare the statement
	stmt, _ := db.Prepare(sqlStr)

	//format all vals at once
	_, resErr := stmt.Exec(vals...)
	fmt.Println(resErr)
	if resErr != nil {
		return resErr
	}

	return nil
}

func CreateUserOrderItems(user string, order_id int64, data []AddItemConfig) error {
	db := database.GetInstance()
	// retrieve groupbuy_id

	sqlStr := "INSERT INTO item(user_id,item,price,quantity, groupbuy_id, order_id) VALUES "
	vals := []interface{}{}
	for _, row := range data {
		sqlStr += "(?,?, ?, ?,?,?),"
		vals = append(vals, user, row.Item, row.Price, row.Quantity, row.Groupbuy_id, order_id)
	}
	//trim the last ,
	sqlStr = strings.TrimSuffix(sqlStr, ",")
	//prepare the statement
	stmt, _ := db.Prepare(sqlStr)

	//format all vals at once
	_, resErr := stmt.Exec(vals...)
	fmt.Println(resErr)
	if resErr != nil {
		fmt.Println(resErr)
		return resErr
	}

	return nil
}
