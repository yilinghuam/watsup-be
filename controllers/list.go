package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"watsup.com/models"
)

// type Env struct {
// 	db *sql.DB
// }

// closure for dependency injection
// func GetLists(env *Env) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		lists, err := models.AllLists(env.db)
// 		if err != nil {
// 			return echo.NewHTTPError(http.StatusBadRequest, "DB query issues")
// 		}
// 		return c.JSON(http.StatusOK, lists)
// 	}
// }

func GetLists(c echo.Context) error {
	lists, err := models.AllLists()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "DB query issues")
	}
	return c.JSON(http.StatusOK, lists)
}

// get one goal from one user
func GetSingleList(c echo.Context) error {
	listName := c.QueryParam("listName")
	lists, err := models.SingleList(listName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "DB query issues")
	}
	return c.JSON(http.StatusOK, lists)
}
