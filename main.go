package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dayptpc00431/momo/configs"
	"github.com/dayptpc00431/momo/middlewares"
	routerv1 "github.com/dayptpc00431/momo/routes/v1"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.New()
	router.NoRoute(noRouteHandle())
	version1 := router.Group("/api/v1")
	routerv1.InitRoutes(version1)
}

func main() {
	port := configs.PORT
	fmt.Println("App running in port: ", port)
	router.Run(":" + port)
}

func noRouteHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			middlewares.CoreMiddleware()(c)
			return
		}
		err := errors.New("Not found handle for the endpoint: " + c.Request.URL.String())
		c.JSON(http.StatusNotFound, gin.H{
			"message:": err,
			"data":     nil,
		})
	}
}
