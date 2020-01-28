package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"
)

const (
	userkey string = "user"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
