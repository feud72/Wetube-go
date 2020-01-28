package routers

import (
	"github.com/feud72/Wetube-go/controllers"
	"github.com/gin-gonic/gin"
)

func VideosRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.VideoIndex)
	router.GET("/:id", controllers.VideoDetail)
	router.GET("/:id/edit", controllers.VideoEdit)
	router.GET("/:id/delete", controllers.VideoDelete)
}
