package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /ping
func r_ping_GET(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
