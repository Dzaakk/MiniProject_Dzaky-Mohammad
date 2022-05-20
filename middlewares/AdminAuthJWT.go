package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const TokenJWT = " "

type Admin struct { //Contoh JSON: {"username": "nama user", "password": "password"}
	Username string `json:"username"`
	Password string `json:"password"`
}

// POST Username and Passowrd Admin "POST -> http://127.0.0.1:1111/api/v1/admin/"
func AdminLoginHandler(c echo.Context) error {
	Admin := Admin{}
	c.Bind(&Admin)
	// username: admin, password: qwerty
	if Admin.Username == "admin" && Admin.Password == "qwerty" {
		token, err := CreateJwtToken(Admin.Username)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Cannot create token")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"success": "ok",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Unauthorized")
}

func CreateJwtToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(TokenJWT))
}
