package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID string `uri:"id" binding:"required"`
}

func UserIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Users index"})
}

func UserDetail(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User Detail", "id": user.ID})
}
