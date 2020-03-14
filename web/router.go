package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/fizzbuzz/stats"
	"github.com/l-lin/fizzbuzz/stats/memory"
)

const (
	fizzBuzzRoute = "/fizz-buzz"
	statsRoute    = "/stats"
)

// NewRouter returns a router with the Logger and Recovery middlewares already attached
func NewRouter() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	// change the repository if we want to use another approach
	s := stats.NewService(memory.NewRepository())

	r := gin.Default()
	r.Use(statsMiddleWare(s))
	r.POST(fizzBuzzRoute, fizzBuzzHandler(s))
	r.GET(statsRoute, statsHandler(s))
	return r
}
