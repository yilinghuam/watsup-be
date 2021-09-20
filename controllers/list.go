package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"watsup.com/auth"
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

func Login(c echo.Context) error {
	// validate token first
	var token auth.TokenConfig
	if err := c.Bind(&token); err != nil {
		fmt.Println(err.Error())
		return err
	}
	email, err := auth.ValidateGoogleToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "token validation error")
	}
	fmt.Println(email)

	// check if email exists in database and create account if does not exist
	user, usererr := models.CheckUserExist(email)
	if usererr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "email search error")
	}
	fmt.Printf("user is %s", user)
	if user == "created" || user == "null" {
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		// This is the information which frontend can use
		// The backend can also decode the token and get admin etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = email
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID          works too)
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"EmailAuth": t,
		})
	}
	if (user != "created" && user != "") || user != "null" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"UserAuth": t,
		})
	}

	return echo.ErrUnauthorized
	// send normal token if user_Id already created
	// return err
}

func DashboardAdd(c echo.Context) error {
	fmt.Println("1")
	fmt.Println(c.Request().Header.Get("user"))

	sentuser := c.Get("user").(*jwt.Token)
	fmt.Print(sentuser)

	claims := sentuser.Claims.(jwt.MapClaims)
	fmt.Print(claims)

	ew := claims["email"].(string)
	fmt.Print(ew)

	var jsonBody DashboardAddConfig
	if err := c.Bind(&jsonBody); err != nil {
		return err
	}
	fmt.Println(jsonBody.Details.Description)
	// get user
	user := c.Request().Header.Get("user")
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
	fmt.Print(user)
	// req := c.Request()
	// headers := req.Header
	// user := headers.Get("user")
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
