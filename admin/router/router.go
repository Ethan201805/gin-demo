package router

import (
	"gin-web-demo/Test/gin-demo/admin/apis"

	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", apis.IndexApi)
	router.POST("/contact", apis.AddQuestions)
	return router
}
