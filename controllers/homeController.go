package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func Join(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Join"})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logout"})
}

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Upload Video"})
}

func EditProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Edit user profile"})
}

func ChangePassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Change user password"})
}
