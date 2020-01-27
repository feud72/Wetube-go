package app

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 미들웨어를 정의하는 방법. 
// 이 방식으로 initialization logic 을 넣을 수 있다.
func DummyMiddleware() gin.HandlerFunc {
	fmt.Println("I'm a dummy!")
	return func(c *gin.Context){
		c.Next()
	}
}

func HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func HandleProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Profile"})
}


func HandleListening() *gin.Engine {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(DummyMiddleware())

	g.GET("/", HandleHome)
	g.GET("/profile", HandleProfile)
	return g
}