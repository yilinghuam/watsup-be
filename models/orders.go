package models

import (
	"fmt"

	"watsup.com/database"
)

func OrdersByUser(user string) ([]GroupbuyDashboard, error) {
	fmt.Println(user)
	rows, err := database.GetInstance().Query(fmt.Sprintf("SELECT watsup.groupbuy.name, watsup.groupbuy.order_date, watsup.order.status FROM watsup.groupbuy INNER JOIN watsup.order ON watsup.order.groupbuy_id=watsup.groupbuy.groupbuy_id WHERE order.user_id = '%s'", user))
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
		err := rows.Scan(&order.Name, &order.Order_date, &order.Status)
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
