package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/ping", func(request *gin.Context) {
		request.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	ginEngine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
