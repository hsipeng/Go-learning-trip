package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter SET UP ROUTERS
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := SetupRouter()
	r.Run()// 0.0.0.0:8080
}
