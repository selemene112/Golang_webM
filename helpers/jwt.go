package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string)string{
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _:= parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken (c *gin.Context) (interface{},error) {
	errResponse := errors.New("Sign in to Proceed")
	headerToken := c.Request.Header.Get("authorization")
	bearer 		:= strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse 

	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token)(interface{},error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errResponse
		}
		return []byte(secretKey),nil
	})
	return token.Claims.(jwt.MapClaims),nil

}

