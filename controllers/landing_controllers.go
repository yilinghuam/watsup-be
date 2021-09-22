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
	// if user is new, send email auth
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
	// send normal auth
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
			"User":     user})
	}

	return echo.ErrUnauthorized
}
