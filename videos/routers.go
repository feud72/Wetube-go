package videos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideosRouter(router *gin.RouterGroup) {
	router.GET("/", VideoIndex)
}

func VideoIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Videos index"})
}
