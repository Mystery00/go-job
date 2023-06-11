package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle(engine *gin.Engine) {
	//注册路由
	engine.GET(apiPath+"/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	engine.GET(externalPath+"/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
