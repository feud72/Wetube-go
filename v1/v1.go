package v1

import (
	"github.com/feud72/Wetube-go/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiBaseEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "API Base Endpoint"})
}

// return 자료형으로 *gin.Engine을 넘겨주어 모듈 분리
func App() *gin.Engine {
	g := gin.Default()

	api := g.Group("/v1")
	routers.HomeRouter(api.Group("/"))
	routers.UsersRouter(api.Group("/users"))
	routers.VideosRouter(api.Group("/videos"))
	return g
}
