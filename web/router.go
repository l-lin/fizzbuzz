package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter returns a router with the Logger and Recovery middlewares already attached
func NewRouter() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/fizz-buzz", fizzBuzzHandler)
	return r
}
