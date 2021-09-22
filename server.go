package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"watsup.com/controllers"
)

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Root route => handler
	e.POST("/login", controllers.Login)

	r := e.Group("/auth")
	config := middleware.JWTConfig{
		SigningKey:    []byte("secret"),
		SigningMethod: "HS256",
	}
	r.Use(middleware.JWTWithConfig(config))
	r.POST("/groupbuy/orders", controllers.AddIndividualOrder)
	r.GET("/groupbuy/:id/view", controllers.ViewOrders)
	r.GET("/groupbuy/:id/orderstatus", controllers.GetOrderStatus)
	r.PATCH("/groupbuy/:id/editorderstatus", controllers.ChangeOrderStatus)
	r.PATCH("/groupbuy/:id/editstatus", controllers.ChangeGroupbuyStatus)
	r.DELETE("/groupbuy/:id/deleteorder", controllers.DeleteOrder)
	r.DELETE("/groupbuy/:id/delete", controllers.DeleteGroupbuy)
	e.GET("/groupbuy/:id", controllers.ViewIndividualGroupbuy)
	e.GET("/groupbuy", controllers.FindGroupbuy)

	r.POST("/dashboard-add", controllers.DashboardAdd)
	r.GET("/dashboard-view", controllers.DashboardView)

	fmt.Println("still working")
	e.Logger.Fatal(e.Start(":8000"))

}
