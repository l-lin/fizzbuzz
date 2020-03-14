package web

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/fizzbuzz/generator"
	"github.com/l-lin/fizzbuzz/stats"
)

func fizzBuzzHandler(s *stats.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var parameters map[string]interface{}
		if err := c.ShouldBindJSON(&parameters); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		int1, int2, limit, str1, str2, err := generator.ToParameters(parameters)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := generator.Generate(int1, int2, limit, str1, str2)
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

func statsMiddleWare(s *stats.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := bytes.NewBuffer(make([]byte, 0))
		r := io.TeeReader(c.Request.Body, b)
		defer c.Request.Body.Close()
		var parameters map[string]interface{}
		if err := json.NewDecoder(r).Decode(&parameters); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		go s.Increment(c.FullPath(), parameters)
		c.Request.Body = ioutil.NopCloser(b)
		c.Next()
	}
}
