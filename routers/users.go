package routers

import (
	"github.com/feud72/Wetube-go/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.UserIndex)
	router.GET("/:id", controllers.UserDetail)
}
