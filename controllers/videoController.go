package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Video struct {
	ID string `uri:"id" binding:"required"`
}

func VideoIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Videos index"})
}

func VideoDetail(c *gin.Context) {
	var video Video
	if err := c.ShouldBindUri(&video); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Videos Detail", "id": video.ID})
}

func VideoEdit(c *gin.Context) {
	var video Video
	if err := c.ShouldBindUri(&video); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Edit Video", "id": video.ID})
}

func VideoDelete(c *gin.Context) {
	var video Video
	if err := c.ShouldBindUri(&video); err != nil {
		c.JSON(400, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete Video", "id": video.ID})
}
