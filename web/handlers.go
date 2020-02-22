package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/fizzbuzz/generator"
	"github.com/l-lin/fizzbuzz/model"
	"github.com/l-lin/fizzbuzz/stats"
)

func fizzBuzzHandler(s *stats.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.Parameters
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		go s.Increment(c.FullPath(), input)
		result, err := generator.Generate(input.Int1, input.Int2, input.Limit, input.Str1, input.Str2)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func statsHandler(s *stats.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		mostUsed := s.GetMostUsed()
		if mostUsed == nil {
			c.JSON(http.StatusOK, gin.H{})
		} else {
			c.JSON(http.StatusOK, *mostUsed)
		}
	}
}
