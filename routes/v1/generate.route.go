package v1

import (
	controllers "github.com/dayptpc00431/Momo-Payment/controllers/v1"
	"github.com/gin-gonic/gin"
)

func SetGenerateURL(router *gin.RouterGroup) {
	router.POST("/generate", controllers.GenerateURL)
}
