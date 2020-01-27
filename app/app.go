package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func HandleProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Profile"})
}

// return 자료형으로 *gin.Engine을 넘겨주어 모듈 분리
func HandleListening() *gin.Engine {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.GET("/", HandleHome)
	g.GET("/profile", HandleProfile)
	return g
}
