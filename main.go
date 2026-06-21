package main

import (
	"Baibai/database"
	"Baibai/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB() // FAil first db connection
	router := gin.Default()
	routes.SetupRouter(router)
	router.Run(":8080")
}
