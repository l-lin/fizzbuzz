package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/fizzbuzz/stats"
	"github.com/rs/zerolog/log"
)

const (
	fizzBuzzRoute = "/fizz-buzz"
	statsRoute    = "/requests/stats"
)

// NewRouter returns a router with the Logger and Recovery middlewares already attached
func NewRouter(repo stats.Repository) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	s := stats.NewService(repo)

	r := gin.Default()
	r.Use(statsMiddleWare(s))
	r.POST(fizzBuzzRoute, fizzBuzzHandler(s))
	r.GET(statsRoute, statsHandler(s))
	// for debugging purpose
	for _, routeInfo := range r.Routes() {
		log.Debug().
			Str("path", routeInfo.Path).
			Str("handler", routeInfo.Handler).
			Str("method", routeInfo.Method).
			Msg("registered routes")
	}
	return r
}
