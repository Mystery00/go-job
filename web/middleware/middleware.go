package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetMiddleware(router *gin.Engine) {
	router.Use(logMiddleware)
	router.Use(recoverMiddleware)
	router.NoRoute(noRouteMiddleware)
}
