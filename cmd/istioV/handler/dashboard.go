package handler

import "github.com/gin-gonic/gin"

// DisplayDashboard display dashboard
func DisplayDashboard(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": "",
		"msg":  "dashboard",
	})
}
