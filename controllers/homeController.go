package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	userkey string = "user"
)

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Home"})
}

func Join(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Join"})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}
	// Check for username and password match
	if username != "hello" || password != "itsme" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
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
