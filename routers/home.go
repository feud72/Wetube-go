package routers

import (
	"github.com/feud72/Wetube-go/controllers"
	"github.com/gin-gonic/gin"
)

func HomeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeIndex)
	router.GET("/join", controllers.Join)
	router.GET("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/search", controllers.Search)
	router.GET("/upload", controllers.Upload)
	router.GET("/edit-profile", controllers.EditProfile)
	router.GET("/change-password", controllers.ChangePassword)
}
