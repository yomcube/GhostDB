package main

import (
	"fmt"
	"io"
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
	if err := os.MkdirAll("rkg", 0777); err != nil { // RKG Folder
		panic(err)
	}
	if err := os.MkdirAll("miigx", 0777); err != nil { // MiiGX Folder
		panic(err)
	}
	if err := os.MkdirAll("mnms", 0777); err != nil { // Cache of Mii Renders
		panic(err)
	}

	ginEngine := gin.Default()

	ginEngine.GET("/ping", func(request *gin.Context) {
		request.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ginEngine.GET("/mnms/:id", func(request *gin.Context) {
		// TODO: handle panics
		id := request.Params.ByName("id")

		path := fmt.Sprintf("./mnms/%s.png", id)
		if _, err := os.Stat(path); err == nil {
			request.File(path)
			return
		}

		miiImgReq, err := http.Get(fmt.Sprintf("https://studio.mii.nintendo.com/miis/image.png?data=%s&type=face&expression=normal&width=512&bgColor=FFFFFF00&clothesColor=default", id))
		if err != nil || miiImgReq.StatusCode != 200 {
			panic(err)
		}
		defer miiImgReq.Body.Close()

		outfile, err := os.Create(path)
		// File closed at the end of the function!
		if err != nil {
			panic(err)
		}

		_, errx := io.Copy(outfile, miiImgReq.Body)
		if errx != nil {
			panic(errx)
		}

		outfile.Close()
		request.File(path)
	})

	ginEngine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
