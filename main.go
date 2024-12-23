package main

import (
	"github.com/yomcube/GhostDB/routes"
	"github.com/yomcube/GhostDB/sql"
	"github.com/yomcube/GhostDB/utils"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup PostgreSQL (Still requires manual creation of the database)
	fmt.Println("Setting up database...")
	cfg, exit_code := sql.SetupConfig()
	if exit_code == 1 {
		fmt.Println("DB config does not exist! Creating default config and exiting...")
		return
	}
	cfg.SetupDatabase()

	// Setup folder structure
	fmt.Println("Ensuring directories...")
	utils.EnsureDirectory("rkg") // RKG Folder
	utils.EnsureDirectory("miigx") // MiiGX Folder
	utils.EnsureDirectory("mnms") // Cache of Mii Renders

	fmt.Println("Initializing router...")
	router := gin.Default()
	routes.InitRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
