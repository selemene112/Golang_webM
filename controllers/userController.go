package controllers

import (
	// "http"
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
	
)

func UserRegister(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON{
		c.ShouldBindJSON(&User)
	
	}else {
		c.ShouldBind(&User)

	}
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": 	User.ID,
		"email": User.Email,
		"username": User.Username,
		"age": User.Age,


	})


}

func UserLogin(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON{
		c.ShouldBindJSON(&User)
	
	}else {
		c.ShouldBind(&User)

	}
	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" :"Unathorized",
			"message": "Invalid Email / Password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password),[]byte(password))

	if !comparePass{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error" :"Unathorized",
			"message": "Invalid Email / Password",

		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK,gin.H{
		"token" : token,
	})

}


