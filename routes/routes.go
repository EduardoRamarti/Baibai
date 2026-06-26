package routes

import (
	"Baibai/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas las rutas de la aplicación en el motor Gin.
// Define tres endpoints:
// - GET /ping: Endpoint de verificación de salud que retorna "pong" para confirmar que el servidor está activo.
// - POST /products: Crea un nuevo producto en la base de datos.
// - GET /products: Obtiene la lista de todos los productos disponibles.
func SetupRouter(r *gin.Engine) {
	// Ruta de verificación de salud (health check)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Ruta POST para crear un nuevo producto
	r.POST("/products", controllers.CreateProduct)

	// Ruta GET para obtener todos los productos
	r.GET("/products", controllers.RetrieveProducts)
}
