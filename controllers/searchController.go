package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchTerm struct {
	Term string `form:"term"`
}

func Search(c *gin.Context) {
	var searchTerm SearchTerm
	if err := c.ShouldBindQuery(&searchTerm); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Search", "term": searchTerm.Term})
}
