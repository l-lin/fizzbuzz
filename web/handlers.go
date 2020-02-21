package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/fizzbuzz/generator"
)

func fizzBuzzHandler(c *gin.Context) {
	var input fizzBuzzInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := generator.Generate(input.Int1, input.Int2, input.Limit, input.Str1, input.Str2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

type fizzBuzzInput struct {
	Int1  int    `json:"int1" binding:"required"`
	Int2  int    `json:"int2" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}
