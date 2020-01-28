package controllers

import (
	"github.com/feud72/Wetube-go/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	var term forms.SearchTerm
	if err := c.ShouldBindQuery(&term); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Search", "term": term.Term})
}
