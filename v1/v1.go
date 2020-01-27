package v1

import (
	"github.com/feud72/Wetube-go/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func HandleProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Profile"})
}

func ApiBaseEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "API Base Endpoint"})
}

// return 자료형으로 *gin.Engine을 넘겨주어 모듈 분리
func BaseEndpoint() *gin.Engine {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.GET("/", HandleHome)
	api := g.Group("/api")
	api.GET("/", ApiBaseEndpoint)
	users.UsersRegister(api.Group("/users"))
	return g
}
