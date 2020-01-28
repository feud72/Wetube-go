package controllers

import (
	"github.com/feud72/Wetube-go/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Users index"})
}

func UserDetail(c *gin.Context) {
	var user forms.UserURI
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User Detail", "id": user.ID})
}
