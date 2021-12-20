package v1

import (
	"fmt"
	"net/http"

	"github.com/dayptpc00431/momo/commons"
	"github.com/dayptpc00431/momo/services"
	"github.com/gin-gonic/gin"
)

func GenerateURL(c *gin.Context) {
	var dto commons.RequestDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		fmt.Println("Can not parse request body to RequestDTO. err=", err)
		return
	}
	fmt.Println("Data body: ", dto)
	data := services.GenerateURL(dto)
	c.JSON(http.StatusOK, gin.H{
		"message": "Generate data success",
		"data":    data,
	})
}
