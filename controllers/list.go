package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"watsup.com/models"
)

type DashboardConfig struct {
	Host []models.GroupbuyDashboard `json:"host"`
	User []models.GroupbuyDashboard `json:"user"`
}
type DashboardDetailsConfig struct {
	Closing_date     string `json:"closing_date"`
	Delivery_options bool   `json:"delivery_options"`
	Delivery_price   int64  `json:"delivery_price"`
	Description      string `json:"description"`
	Name             string `json:"name"`
	Order_date       string `json:"order_date"`
}
type DashboardSetupConfig struct {
	Item     string `json:"item"`
	Price    int8   `json:"price"`
	Quantity int16  `json:"quantity"`
}
type DashboardAddConfig struct {
	Details DashboardDetailsConfig `json:"Details"`
	Setup   []DashboardSetupConfig `json:"Setup"`
}

func DashboardAdd(c echo.Context) error {
	fmt.Println("1")
	fmt.Println(c.Request().Body)
	var jsonBody DashboardAddConfig
	if err := c.Bind(&jsonBody); err != nil {
		return err
	}
	fmt.Println(jsonBody.Details.Description)
	// get user
	user := c.Request().Header.Get("user")

	// optionsbool:= jsonBody.Details.Delivery_options

	// if optionsbool {
	// 	price, err := strconv.ParseInt(c.FormValue("delivery_price"), 10, 64)
	// 	table := models.CreateGroupbuyConfig{
	// 		User_id:          c.Request().Header.Get("user"),
	// 		Name:             c.FormValue("name"),
	// 		Description:      c.FormValue("description"),
	// 		Order_date:       c.FormValue("user_id"),
	// 		Closing_date:     "01-12-2021",
	// 		Delivery_options: optionsbool,
	// 		Delivery_price:   price,
	// 	}
	// } else {
	// 	table := models.CreateGroupbuyConfig{
	// 		User_id:          "ling",
	// 		Name:             "eat 3 bowls",
	// 		Description:      "great food",
	// 		Order_date:       "12-12-2021",
	// 		Closing_date:     "01-12-2021",
	// 		Delivery_options: true,
	// 		Delivery_price:   5,
	// 	}
	// }
	table := models.CreateGroupbuyConfig{
		User_id:          user,
		Name:             jsonBody.Details.Name,
		Description:      jsonBody.Details.Description,
		Order_date:       jsonBody.Details.Order_date,
		Closing_date:     jsonBody.Details.Closing_date,
		Delivery_options: jsonBody.Details.Delivery_options,
		Delivery_price:   jsonBody.Details.Delivery_price,
	}
	err := models.CreateGroupbuy(table)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "create groupbuy issues")
	}

	// create items using formsetup
	return err
}

func DashboardView(c echo.Context) error {
	req := c.Request()
	headers := req.Header
	user := headers.Get("user")
	lists, err := models.GroupbuyByHost(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "host groupbuy query issues")
	}
	orders, err := models.OrdersByUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ordered groupbuy query issues")
	}

	var dashboard DashboardConfig
	dashboard.Host = lists
	dashboard.User = orders

	return c.JSON(http.StatusOK, dashboard)
}
