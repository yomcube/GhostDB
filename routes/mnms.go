package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GET /mnms/:id
func r_mnms_id_GET(ctx *gin.Context) {
	// TODO: handle panics
	id := ctx.Params.ByName("id")

	path := fmt.Sprintf("./mnms/%s.png", id)
	if _, err := os.Stat(path); err == nil {
		ctx.File(path)
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
	ctx.File(path)
}
