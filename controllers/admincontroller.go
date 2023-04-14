package controllers

// import (
// 	"encoding/json"
// 	"final/database"
// 	"final/models"
// 	"net/http"

// 	"strconv"

// 	// "github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	// "gorm.io/gorm"

// 	// "jwt_token/database"
// 	"final/helpers"
// 	// "jwt_token/models"
// 	// "net/http"
// 	// "strconv"

// 	"github.com/dgrijalva/jwt-go"
// 	// "github.com/gin-gonic/gin"
// )

// //=============================================================Admin ALL SHOW Product==========================================
// func AdminIndex(c *gin.Context){
// 	db := database.GetDB()
// 	var Product []models.Product
// 	db.Find(&Product)

// 	c.JSON(http.StatusOK, gin.H{"Product": Product})

// }

//============================================================End All SHOW PRODUCT============================================

//===========================================================UPDATE PRODUCT==================================================

// func UpdateProduct1(c *gin.Context){
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

// //===========================================================END UPDATE PRODUCT=============================================

// //==============================================================Delete=====================================================
// func Delete(c *gin.Context){
// 	db := database.GetDB()

// 	var Product models.Product
// 	var input struct{
// 		Id json.Number

// 	}
// 	// input := map[string]string{"Id":"0"}

// 	//ERROR data
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": err.Error()})
// 		return
// 	}
// 	Id, _ := input.Id.Int64()
// 	//strconv.ParseInt(input["Id"], 10, 64)
// 	if db.Delete(&Product, Id).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": "Tindak Dapat Menghapus"})
// 		return
// 	}
// 	//END ERROR DATA

// 	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})//data Berhasil

// }

// //=============================================================END DELETE==================================================


