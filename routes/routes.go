package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/ping", r_ping_GET)
	router.GET("/mnms/:id", r_mnms_id_GET)
}
