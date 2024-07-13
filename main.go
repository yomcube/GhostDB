package main

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yomcube/GhostDB/sql"
)

func main() {
	// Setup PostgreSQL (Still requires manual creation of the database)
	cfg, exit_code := sql.SetupConfig()
	if exit_code == 1 {
		return
	}
	cfg.SetupDatabase()

	// Setup folder structure
	// Using MkdirAll because it doesn't error out if the folder exists
	if err := os.MkdirAll("rkg", fs.ModeDir); err != nil { // RKG Folder
		panic(err)
	}
	if err := os.MkdirAll("miigx", fs.ModeDir); err != nil { // MiiGX Folder
		panic(err)
	}
	if err := os.MkdirAll("mnms", fs.ModeDir); err != nil { // Cache of Mii Renders
		panic(err)
	}

	ginEngine := gin.Default()
	ginEngine.GET("/ping", func(request *gin.Context) {
		request.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	ginEngine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
