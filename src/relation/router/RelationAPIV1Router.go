package router

import "github.com/gin-gonic/gin"

func LoadApiV1Router(r *gin.Engine) {
	v1 := r.Group("/douyin/relation")
	v1.GET("hello", HelloWorld)
}
