package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/fizzbuzz/stats"
	"github.com/l-lin/fizzbuzz/stats/memory"
)

const (
	fizzBuzzRoute = "/fizz-buzz"
	statsRoute    = "/requests/stats"
)

var repositories = map[stats.Mode]stats.Repository{
	stats.Memory: memory.NewRepository(),
}

// NewRouter returns a router with the Logger and Recovery middlewares already attached
func NewRouter(statsStorageMode string) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	m := stats.GetMode(statsStorageMode)
	s := stats.NewService(repositories[m])

	r := gin.Default()
	r.Use(statsMiddleWare(s))
	r.POST(fizzBuzzRoute, fizzBuzzHandler(s))
	r.GET(statsRoute, statsHandler(s))
	return r
}
