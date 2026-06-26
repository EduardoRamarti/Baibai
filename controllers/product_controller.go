package controllers

import (
	"Baibai/database"
	"Baibai/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct recibe una petición POST con un JSON que contiene los datos de un producto.
// Valida el JSON, genera un ID único usando UUID, lo guarda en la base de datos PostgreSQL
// y devuelve el objeto creado con estado HTTP 201 (Created) o un error si ocurre un problema.
func CreateProduct(c *gin.Context) {
	var p models.Product

	// ShouldBindJSON valida el cuerpo JSON de la petición y lo mapea al modelo Product.
	// Si hay un error de validación (JSON inválido o campos faltantes), retorna un error 400 Bad Request.
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Genera un identificador único (UUID v4) para el nuevo producto.
	p.ID = uuid.NewString()

	// Ejecuta una consulta INSERT para guardar el producto en la tabla 'products' de PostgreSQL.
	// Utiliza parámetros preparados ($1, $2, $3, $4) para prevenir inyección SQL.
	_, erro := database.DB.Exec("INSERT INTO products (id, productname, stock, price) VALUES ($1, $2, $3, $4)", p.ID, p.ProductName, p.Stock, p.Price)

	// Si la inserción falla, registra el error en los logs y retorna un error 500 Internal Server Error.
	if erro != nil {
		log.Printf("error: %v", erro)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "It was impossible to save product",
		})
		return
	}

	// Si todo es exitoso, retorna el producto creado con estado HTTP 201 (Created).
	c.JSON(http.StatusCreated, p)
}

// RetrieveProducts maneja las peticiones GET para obtener todos los productos.
// Ejecuta una consulta SELECT sobre la tabla 'products', itera sobre los resultados,
// escanea cada fila en un modelo Product y retorna la lista completa de productos.
func RetrieveProducts(c *gin.Context) {
	var listOfProducts []models.Product

	// Ejecuta una consulta SELECT para obtener todos los registros de la tabla 'products'.
	rows, err := database.DB.Query("SELECT * FROM products")
	// Si ocurre un error en la consulta, registra el error y retorna un error 503 Service Unavailable.
	if err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "It was impossible to retrieve products",
		})
		return
	}

	// Asegura que el objeto rows se cierre después de procesar todos los resultados.
	defer rows.Close()
	// Itera sobre cada fila obtenida de la consulta SELECT.
	for rows.Next() {
		var p models.Product

		// Escanea los valores de la fila actual y los asigna a los campos del modelo Product.
		// Si ocurre un error durante el escaneo, registra el error y retorna un error 503.
		if err := rows.Scan(&p.ID, &p.ProductName, &p.Stock, &p.Price); err != nil {
			log.Printf("Error: %v", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "It was impossible to retrieve products",
			})
			return
		}

		// Agrega el producto procesado a la lista.
		listOfProducts = append(listOfProducts, p)
	}

	// Retorna la lista completa de productos con estado HTTP 200 (OK) en formato JSON.
	c.JSON(http.StatusOK, listOfProducts)
}
