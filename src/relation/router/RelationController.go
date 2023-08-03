package router

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /douyin/relation

// HelloWorld godoc
// @Summary HelloWorld example
// @Tags test
// @Description do test
// @Accept json
// @Success 200 {string} string "{"msg": "hello world"}"
// @Router /hello [get]
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello world"})
}
