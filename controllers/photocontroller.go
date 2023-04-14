package controllers

import (
	"encoding/json"
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"

	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//===================================================================Create==============================================
func CreatePhoto(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON{
		c.ShouldBindJSON(&Photo)
	} else{
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "BAD REQ",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
	

}

//=====================================================================END CREATE=========================================
//===================================================================GET ALL =============================================
func PhotoIndex(c *gin.Context){
	db := database.GetDB()
	var Photo []models.Photo
	db.Find(&Photo)

	c.JSON(http.StatusOK, gin.H{"Product": Photo})

}
//==================================================================END GET ALL============================================

//==================================================================GET BY ID =============================================
func PhotoIDIndex(c *gin.Context){
	
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
	
			// userID := uint(userData["userid"].(float64))
		// 	userID := uint(userData["userid"].(float64))
	
		Product.UserID = userID
		Product.ID = uint(productId)
	
		// err := db.Model(&Product).Where("id = ? ", userID).Find(models.Photo{}).Error
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

//===================================================================END GET BY ID=========================================

//===================================================================PUT ===================================================
func UpdatePhoto(c *gin.Context){
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

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ? ", productId).Updates(models.Photo{Title: Product.Title, Caption: Product.Caption, Photo_Url: Product.Photo_Url    }).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad Request",
			"message": err.Error(),

		})
		return
	}

	c.JSON(http.StatusOK, Product)
}


//====================================================================END PUT================================================


//=====================================================================DELETE================================================

func Delete(c *gin.Context){
	db := database.GetDB()

	var Product models.Photo
	var input struct{
		Id json.Number

	}
	// input := map[string]string{"Id":"0"}

	//ERROR data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": err.Error()})
		return
	}
	Id, _ := input.Id.Int64()
	//strconv.ParseInt(input["Id"], 10, 64)
	if db.Delete(&Product, Id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": "Tindak Dapat Menghapus"})
		return
	}
	//END ERROR DATA

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})//data Berhasil

}



//======================================================================END DELETE==========================================