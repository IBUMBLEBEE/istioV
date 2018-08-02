package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pod(c *gin.Context) {
	c.String(http.StatusOK, "pod list")
}

func main() {
	router := gin.Default()

	podlist := router.Group("/pod")
	podlist.GET("/list", pod)

	router.Run(":9090")
}
