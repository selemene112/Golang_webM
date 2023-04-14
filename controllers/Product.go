package controllers

import (
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func CreateProduct(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON{
		c.ShouldBindJSON(&Product)
	} else{
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "BAD REQ",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Product)
	

}

// func UpdateProduct(c *gin.Context){
// 	db := database.GetDB()
// 	userData := c.MustGet("userData").(jwt.MapClaims)
// 	contentType := helpers.GetContentType(c)
// 	Product := models.Product{}

// 	productId, _ := strconv.Atoi(c.Param("productId"))
// 	userID := uint(userData["id"].(float64))


// 	if contentType == appJSON{
// 		c.ShouldBindJSON(&Product)

// 	} else {
// 		c.ShouldBind(&Product)
// 	}

// 	Product.UserID = userID
// 	Product.ID = uint(productId)

// 	err := db.Model(&Product).Where("id = ? ", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description  }).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "bad Request",
// 			"message": err.Error(),

// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, Product)
// }

// ==============================================================GET Product===========================================
func UserIndex(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Photo{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))


	if contentType == appJSON{
		c.ShouldBindJSON(&Product)

	} else {
		c.ShouldBind(&Product)
	}

	// 	userID := uint(userData["userid"].(float64))
	// 	userID := uint(userData["userid"].(float64))

	Product.UserID = userID
	Product.User.ID = uint(productId)

	// err := db.Model(&Product).Where("id = ? ", userID).Find(models.Product{}).Error
	// err := db.Select("user_id").Find(&Product, uint(productId)).Error
	err := db.Debug().Find(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad Request",
			"message": err.Error(),

		})
		return
	}


	c.JSON(http.StatusOK, Product)
}

//===============================================================END GET==============================================


func AdminIndex1(c *gin.Context){
	db := database.GetDB()
	var Product []models.Photo
	db.Find(&Product)

	c.JSON(http.StatusOK, gin.H{"Product": Product})

}

//==========================================================================

// func UserIndex1(c *gin.Context){
// 	db := database.GetDB()
// 	userData := c.MustGet("userData").(jwt.MapClaims)
// 	contentType := helpers.GetContentType(c)
// 	var Product models.Product

// 	// Product := []models.Product
// 	userID := uint(userData["userid"].(float64))

// 	if contentType == appJSON{
// 		c.ShouldBindJSON(&Product)
// 	} else{
// 		c.ShouldBind(&Product)
// 	}

// 	Product.UserID = userID

// 	// err := db.Debug().Find(&Product).Error
// 	// err :=db.Debug().Model(&Product).Where("UserId = ?", userID).Find(&Product).Error
// 	err := db.Model(&Product).Where(" = ? ", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description  }).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"err": "BAD REQ",
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	// c.JSON(http.StatusCreated, Product)
// 	c.JSON(http.StatusOK, gin.H{"Product": Product})
	

// }


//=================================================================================coba==================================

