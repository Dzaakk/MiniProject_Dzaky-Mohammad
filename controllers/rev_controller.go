package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET All Review Data "GET -> http://127.0.0.1:1111/api/v1/reviews/"
func GETAllreviews(c echo.Context) error {
	var reviews []models.Reviews
	if err := config.DB.Find(&reviews).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, reviews)
}

// CREATE new Review "POST -> http://127.0.0.1:1111/api/v1/reviews/"
// {
//		"userid": ,
//		"restid": ,
//		"comment" : "",
//		"rating" :
// }
func CREATEReview(c echo.Context) error {

	review := models.Reviews{}
	if err := c.Bind(&review); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	fmt.Printf("Before insert: %#v\n", review)
	if err := config.DB.Save(&review).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, review)
}
