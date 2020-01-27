package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.GET("/", UsersRegistration)
}

func UsersRegistration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Users Registration"})
}
