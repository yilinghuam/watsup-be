package models

import (
	"fmt"

	"watsup.com/database"
)

type OrderConfig struct {
	Groupbuy_id int16
	User_id     string
	Address     string
}

func ChangeOrderStatusByOrderID(status string, order_id int) error {
	db := database.GetInstance()
	_, err := db.Exec("UPDATE watsup.order set status=? where order_id=? ", status, order_id)
	if err != nil {
		fmt.Println("update order status issue")
		return err
	}
	return nil
}

func GetStatusByGroupbuy(groupbuy_id int) ([]string, error) {
	rows, err := database.GetInstance().Query(fmt.Sprintf("SELECT status FROM watsup.order where groupbuy_id ='%d'", groupbuy_id))
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return nil, err
	}
	defer rows.Close()

	var orders []string

	//for loop to read through each row of data
	for rows.Next() {
		var order string
		err := rows.Scan(&order)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return nil, err
		}

		orders = append(orders, order)

	}
	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	return orders, nil
}
func DeleteOrdersByGroupbuy(groupbuy_id int) error {
	_, err := database.GetInstance().Exec("DELETE FROM watsup.order where groupbuy_id =?", groupbuy_id)
	if err != nil {
		fmt.Println("delete orders issue")
		fmt.Println(err)
		return err
	}

	return nil
}
func DeleteOrdersByOrderId(order_id int) error {
	_, err := database.GetInstance().Exec("DELETE FROM watsup.order where order_id =?", order_id)
	if err != nil {
		fmt.Println("delete orders issue")
		fmt.Println(err)
		return err
	}

	return nil
}

func CreateOrderByUser(groupbuy_id int16, user_id string, address string) (int64, error) {
	db := database.GetInstance()
	// retrieve groupbuy_id
	fmt.Println(1)
	res, err := db.Exec(fmt.Sprintf("INSERT INTO watsup.order(groupbuy_id,user_id,address)VALUES(%d,'%s','%s')", groupbuy_id, user_id, address))
	fmt.Println(res)
	fmt.Println(2)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()

	fmt.Println(id)
	// close if there is an error
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func OrdersByUser(user string) ([]GroupbuyDashboard, error) {
	fmt.Println(user)
	rows, err := database.GetInstance().Query(fmt.Sprintf("SELECT watsup.groupbuy.name, watsup.groupbuy.order_date, watsup.order.status, watsup.order.groupbuy_id FROM watsup.groupbuy INNER JOIN watsup.order ON watsup.order.groupbuy_id=watsup.groupbuy.groupbuy_id WHERE order.user_id = '%s'", user))
	// close if there is an error
	if err != nil {
		fmt.Println("query issue")
		return nil, err
	}
	defer rows.Close()

	var orders []GroupbuyDashboard

	//for loop to read through each row of data
	for rows.Next() {
		var order GroupbuyDashboard
		err := rows.Scan(&order.Name, &order.Order_date, &order.Status, &order.Groupbuy_id)
		if err != nil {
			fmt.Println("scan issue")
			fmt.Println(err)
			return nil, err
		}

		orders = append(orders, order)

	}
	if err = rows.Err(); err != nil {
		fmt.Println("iteration error")
		return nil, err
	}
	return orders, nil
}
