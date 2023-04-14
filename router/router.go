package router

import (
	"final/controllers"
	"final/middlewares"

	"github.com/gin-gonic/gin"
)



func StartApp() *gin.Engine{
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

	
	}

	PhotoRoute := r.Group("/photos")
	{
		PhotoRoute.Use(middlewares.Authentication())
		PhotoRoute.POST("/", controllers.CreatePhoto)
		PhotoRoute.GET("/",controllers.PhotoIndex)
		PhotoRoute.GET("/:productId", middlewares.ProductAuthorization(),controllers.PhotoIDIndex)
		PhotoRoute.PUT("/:productId", middlewares.ProductAuthorization(),controllers.UpdatePhoto)
		PhotoRoute.DELETE("/",controllers.Delete)
		

	}

	SM := r.Group("/sosial")
	{
		SM.Use(middlewares.Authentication())
		SM.POST("/", controllers.CreateSM)
		SM.GET("/",controllers.SMIndex)
		SM.GET("/:productId", middlewares.SMAuthorization(),controllers.SMIndexId)
		SM.PUT("/:productId", middlewares.SMAuthorization(),controllers.UpdateSM)
		SM.DELETE("/",controllers.DeleteSM)
		

		

	}

	COM := r.Group("/comment")
	{
		COM.Use(middlewares.Authentication())
		COM.POST("/", controllers.CreateCOM)
		COM.GET("/",controllers.COMIndex)
		COM.GET("/:productId", middlewares.COMAuthorization(),controllers.COMIDIndex)
		COM.PUT("/:productId", middlewares.COMAuthorization(),controllers.UpdateCOM)
		COM.DELETE("/",controllers.DeleteCOM)
		
		

	}


	
	return r
}