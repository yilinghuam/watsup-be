package controllers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
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

type DashboardAddConfig struct {
	Details DashboardDetailsConfig        `json:"Details"`
	Setup   []models.DashboardSetupConfig `json:"Setup"`
}

// dashboard functions
func DashboardAdd(c echo.Context) error {
	fmt.Println("1")
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	user := claims["name"].(string)
	fmt.Print(user)

	var jsonBody DashboardAddConfig
	if err := c.Bind(&jsonBody); err != nil {
		return err
	}
	fmt.Println(jsonBody.Details.Description)
	// get user
	// set table for groupbuy creation
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
	name := jsonBody.Details.Name
	setup := jsonBody.Setup

	Itemerr := models.CreateGroupbuyItems(user, name, setup)
	if Itemerr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "create groupbuy items issues")
	}

	return err
}

func DashboardView(c echo.Context) error {
	fmt.Println("1")
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	user := claims["name"].(string)

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
