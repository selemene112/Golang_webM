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

//==============================================================Create =====================================================
func CreateSM(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.SosialMedias{}
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

//==============================================================END CREATE==================================================

//=============================================================GET ALL=====================================================
func SMIndex(c *gin.Context){
	db := database.GetDB()
	var Photo []models.SosialMedias
	db.Find(&Photo)

	c.JSON(http.StatusOK, gin.H{"Sosial": Photo})

}

//==============================================================END GET ALL================================================




//================================================================GET BY ID================================================
func SMIndexId(c *gin.Context){
	
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.SosialMedias{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))


	if contentType == appJSON{
		c.ShouldBindJSON(&Product)

	} else {
		c.ShouldBind(&Product)
	}


	Product.UserID = userID
	Product.ID = uint(productId)
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

//================================================================END GET BY ID============================================


//==================================================================PUT====================================================
func UpdateSM(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.SosialMedias{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON{
		c.ShouldBindJSON(&Product)

	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ? ", productId).Updates(models.SosialMedias{Name: Product.Name, SosialMedia_Url: Product.SosialMedia_Url}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad Request",
			"message": err.Error(),

		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

//=================================================================END PUT================================================


//==================================================================DELETE================================================

func DeleteSM(c *gin.Context){
	db := database.GetDB()

	var Product models.SosialMedias
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


//===================================================================END DELETE=======================================
