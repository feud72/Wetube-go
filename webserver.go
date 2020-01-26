package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func handleProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Profile"})
}


func main() {
	addr := ":3000"
	g := gin.Default()
	g.GET("/", handleHome)
	g.GET("/profile", handleProfile)
	log.Fatal(http.ListenAndServe(addr, g))
}