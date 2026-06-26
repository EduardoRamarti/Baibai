package main

import (
	"Baibai/database"
	"Baibai/routes"

	"github.com/gin-gonic/gin"
)

// main es el punto de entrada de la aplicación.
// Realiza los siguientes pasos:
// 1. Conecta a la base de datos PostgreSQL.
// 2. Inicializa el motor web Gin con configuración predeterminada.
// 3. Configura todas las rutas de la API.
// 4. Inicia el servidor en el puerto 8080.
func main() {
	// Establece la conexión con la base de datos PostgreSQL.
	// Si la conexión falla, la aplicación se detiene inmediatamente.
	database.ConnectDB()

	// Crea una instancia del router Gin con configuración predeterminada.
	// Incluye logging y recuperación de panics automáticos.
	router := gin.Default()

	// Registra todas las rutas de la API en el router.
	routes.SetupRouter(router)

	// Inicia el servidor HTTP en el puerto 8080.
	// Permanece en ejecución escuchando peticiones entrantes.
	router.Run(":8080")
}
