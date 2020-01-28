package v1

import (
	"github.com/feud72/Wetube-go/routers"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiBaseEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "API Base Endpoint"})
}

func Engine() *gin.Engine {
	g := gin.Default()

	api := g.Group("/v1")
	api.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	routers.HomeRouter(api.Group("/"))
	routers.UsersRouter(api.Group("/users"))
	routers.VideosRouter(api.Group("/videos"))
	return g
}
