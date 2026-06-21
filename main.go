package main

import (
	"Baibai/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRouter(router)
	router.Run(":8080")
}
