package main

import (
	"log"
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

func handleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func handleProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Profile"})
}


func main() {
	addr := ":3000"
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(DummyMiddleware())

	g.GET("/", handleHome)
	g.GET("/profile", handleProfile)
	log.Fatal(http.ListenAndServe(addr, g))
}