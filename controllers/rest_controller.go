package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//get data all restaurant "GET -> http://127.0.0.1:1111/api/v1/restaurants/"
func GETAllrestaurant(c echo.Context) error {
	var restaurants []models.Restaurants
	if err := config.DB.Find(&restaurants).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, restaurants)
}

//get data spesific restaurant "GET -> http://127.0.0.1:1111/api/v1/restaurants/:RID"
func GETSpesRestaurant(c echo.Context) error {

	restId, err := strconv.Atoi(c.Param("RID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid restaurant id!")
	}
	var restaurant models.Restaurants

	if err := config.DB.First(&restaurant, restId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")

	}
	if restaurant.ID == 0 {
		return c.String(http.StatusNotFound, "restaurant not found")

	}
	return c.JSON(http.StatusOK, restaurant)
}

// create new restaurant "POST -> http://127.0.0.1:1111/api/v1/restaurants/" (AUTHOR ADMIN)
// {
//     "restname" : "",
//     "Cuisine" : "",
//     "address": "",
//     "contact": ""
// }
func CREATERestaurant(c echo.Context) error {

	restaurant := models.Restaurants{}
	if err := c.Bind(&restaurant); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	fmt.Printf("Before insert: %#v\n", restaurant)
	if err := config.DB.Save(&restaurant).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, restaurant)
}

// EDIT Restaurant "PUT -> http://127.0.0.1:1111/api/v1/restaurants/:RID"
func EDITRestaurant(c echo.Context) error {

	restId, err := strconv.Atoi(c.Param("RID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid restaurant id!")
	}
	fmt.Println("Isi restaurantId", restId)
	var restaurant models.Restaurants
	fmt.Printf("Isi restaurant sebelum select %#v\n", restaurant)
	if err := config.DB.First(&restaurant, restId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if restaurant.ID == 0 {
		return c.String(http.StatusNotFound, "restaurant not found")
	}

	fmt.Printf("isi restaurant setelah select %#v\n", restaurant)
	if err := c.Bind(&restaurant); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	fmt.Printf("Isi restaurant setelah bind %#v\n", restaurant)
	fmt.Printf("Before Update : %#v\n", restaurant)
	if err := config.DB.Save(&restaurant).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, restaurant)
}

// DELETE restaurant "DELETE -> http://127.0.0.1:1111/api/v1/restaurants/:RID"
func DELETErestaurant(c echo.Context) error {

	restaurantId, err := strconv.Atoi(c.Param("RID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid restaurant id!")
	}

	var restaurant models.Restaurants
	if err := config.DB.First(&restaurant, restaurantId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if restaurant.ID == 0 {
		return c.String(http.StatusNotFound, "restaurant not found")
	}
	if err := config.DB.Delete(&restaurant).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":           "Deleted",
		"ID":               restaurantId,
		"Restaurant Name ": restaurant.RestName,
	})
}
