package routers

import (
	"github.com/feud72/Wetube-go/controllers"
	"github.com/feud72/Wetube-go/middlewares"
	"github.com/gin-gonic/gin"
)

func HomeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeIndex)
	router.GET("/join", controllers.Join)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/search", controllers.Search)
	router.GET("/upload", controllers.Upload)
	router.Use(middlewares.AuthRequired)
	{
		router.GET("/edit-profile", controllers.EditProfile)
		router.GET("/change-password", controllers.ChangePassword)
	}
}
