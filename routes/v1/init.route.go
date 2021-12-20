package v1

import (
	middlewares "github.com/dayptpc00431/momo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
	// g.Use(loghttp.Middleware())
	g.Use(middlewares.CoreMiddleware())
	SetGenerateURL(g)
}
