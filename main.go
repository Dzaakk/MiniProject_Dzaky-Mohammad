package main

import (
	"MiniProject/config"
	"MiniProject/routes"
)

func main() {

	config.InitDB()
	e := routes.New()
	e.Start(":1111")
}

// package main

// import (
// 	"net/http"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// const secretJwt = "d6ca6c89-ce38-49d4-8d8e-db0c553b4f25" // kunci untuk buat JWT

// func main() {
// 	e := echo.New()

// 	e.GET("/", HelloHandler)                                           // Selalu bisa diakses (tidak butuh authentication)
// 	e.POST("/login", LoginHandler)                                     // link untuk login
// 	e.GET("/secret", SecretHandler, middleware.JWT([]byte(secretJwt))) // Harus via authentication (login dulu)

// 	// eAuth := e.Group("/auth")
// 	// eAuth.Use(middleware.JWT([]byte(secretJwt)))
// 	// eAuth.GET("/secret", SecretHandler) // Harus via authentication (login dulu) /auth/secret
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// func HelloHandler(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

// func CreateJwtToken(username string) (string, error) {
// 	claims := jwt.MapClaims{}
// 	claims["username"] = username
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secretJwt))
// }

// type LoginData struct { //Contoh JSON: {"username": "nama user", "password": "password"}
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// func LoginHandler(c echo.Context) error {
// 	loginData := LoginData{}
// 	c.Bind(&loginData)
// 	// username: admin, password: rahasia
// 	if loginData.Username == "admin" && loginData.Password == "rahasia" {
// 		token, err := CreateJwtToken(loginData.Username)
// 		if err != nil {
// 			return c.String(http.StatusInternalServerError, "Cannot create token")
// 		}
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"success": "ok",
// 			"token":   token,
// 		})
// 	}
// 	return c.String(http.StatusUnauthorized, "Unauthorized")
// }

// func SecretHandler(c echo.Context) error {
// 	token := c.Get("user").(*jwt.Token)
// 	claims := token.Claims.(jwt.MapClaims)
// 	username := claims["username"].(string)
// 	// select * from order where id=username
// 	return c.String(http.StatusOK, "Hello, "+username)
// }
