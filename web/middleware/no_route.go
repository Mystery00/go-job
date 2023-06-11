package middleware

import (
	"github.com/gin-gonic/gin"
)

var noRouteMiddleware gin.HandlerFunc = func(context *gin.Context) {
	panic(`404 Not Found`)
}
