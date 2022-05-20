package routes

import (
	"MiniProject/controllers"
	"MiniProject/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	const TokenJWT = " "

	//users
	e.GET("/api/v1/users/", controllers.GETAllusers, middleware.JWT([]byte(TokenJWT)))
	e.GET("/api/v1/users/:UID", controllers.GETSpesUser)
	e.POST("/api/v1/users/", controllers.CREATEUser)
	e.PUT("/api/v1/users/:UID", controllers.EDITUser, middleware.JWT([]byte(TokenJWT)))
	e.DELETE("/api/v1/users/:UID", controllers.DELETEUser, middleware.JWT([]byte(TokenJWT)))

	//restaurants
	e.GET("/api/v1/restaurants/", controllers.GETAllrestaurant)
	e.GET("/api/v1/restaurants/:RID", controllers.GETSpesRestaurant)
	e.POST("/api/v1/restaurants/", controllers.CREATERestaurant)
	e.PUT("/api/v1/restaurants/:RID", controllers.EDITRestaurant, middleware.JWT([]byte(TokenJWT)))
	e.DELETE("/api/v1/restaurants/:RID", controllers.DELETErestaurant, middleware.JWT([]byte(TokenJWT)))

	//reviews
	e.GET("/api/v1/reviews/", controllers.GETAllreviews)
	e.POST("/api/v1/reviews/", controllers.CREATEReview)

	//admin
	e.POST("/api/v1/admin/", middlewares.AdminLoginHandler)

	return e
}
