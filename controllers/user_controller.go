package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET All User Data "GET -> http://127.0.0.1:1111/api/v1/users/"
func GETAllusers(c echo.Context) error {
	var users []models.Users
	if err := config.DB.Find(&users).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, users)
}

// GET Spesific Data User "GET -> http://127.0.0.1:1111/api/v1/users/:UID"
func GETSpesUser(c echo.Context) error {

	userId, err := strconv.Atoi(c.Param("UID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid user id!")
	}
	var user models.Users

	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")

	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":    userId,
		"Name":  user.Name,
		"Email": user.Email,
	})
}

// CREATE New User "POST -> http://127.0.0.1:1111/api/v1/users/"
// {
// 		"Username": "",
//     	"Password": "",
//     	"Name": "",
//     	"Email": ""
// }
func CREATEUser(c echo.Context) error {

	user := models.Users{}
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	fmt.Printf("Before insert: %#v\n", user)
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, user)
}

// EDIT Spesific User Data "PUT -> http://127.0.0.1:1111/api/v1/users/:UID"
func EDITUser(c echo.Context) error {

	userId, err := strconv.Atoi(c.Param("UID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid user id!")
	}
	fmt.Println("Isi userId", userId)
	var user models.Users
	fmt.Printf("Isi user sebelum select %#v\n", user)
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}

	fmt.Printf("isi user setelah select %#v\n", user)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error2")
	}

	fmt.Printf("Isi user setelah bind %#v\n", user)
	fmt.Printf("Before Update : %#v\n", user)
	if err := config.DB.Save(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error3") //?
	}

	return c.JSON(http.StatusOK, user)
}

// DELETE user "DELETE -> http://127.0.0.1:1111/api/v1/users/:UID"
func DELETEUser(c echo.Context) error {

	userId, err := strconv.Atoi(c.Param("UID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid user id!")
	}

	var user models.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := config.DB.Delete(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error!")
	}
	return c.JSON(http.StatusOK, user)
}
