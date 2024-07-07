package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yomcube/GhostDB/sql"
)

func main() {
	cfg, exit_code := sql.SetupConfig()
	if exit_code == 1 {
		return
	}
	cfg.SetupDatabase()

	fmt.Println(cfg)

	ginEngine := gin.Default()
	ginEngine.GET("/ping", func(request *gin.Context) {
		request.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	ginEngine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
