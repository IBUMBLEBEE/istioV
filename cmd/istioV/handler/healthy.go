package handler

import "github.com/gin-gonic/gin"

// Healthy health check
func Healthy(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": "",
		"msg":  "health check",
	})
}
