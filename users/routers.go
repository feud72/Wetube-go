package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersRouter(router *gin.RouterGroup) {
	router.GET("/", UserIndex)
	router.GET("/edit", UserEdit)
	router.GET("/password", UserPassword)
}

func UserIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Users index"})
}

func UserEdit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User edit"})
}

func UserPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User password"})
}
