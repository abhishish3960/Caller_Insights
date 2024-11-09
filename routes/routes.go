package routes

import (
	"myapp/controllers"
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(utils.JWTMiddleware()) // Use JWT Middleware for authentication

	authorized.GET("/profile", controllers.GetProfile)
	authorized.PUT("/profile", controllers.UpdateProfile)
	authorized.POST("/spam", controllers.MarkAsSpam)
	authorized.GET("/search/name", controllers.SearchByName)
	authorized.GET("/search/phone", controllers.SearchByPhoneNumber)

	return r
}
