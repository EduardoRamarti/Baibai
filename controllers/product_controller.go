package controllers

import (
	"Baibai/database"
	"Baibai/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct recibe un JSON con los datos de un producto, lo guarda en la base de datos y devuelve el objeto creado.
func CreateProduct(c *gin.Context) {
	var p models.Product

	// Valida y enlaza el cuerpo de la petición JSON al modelo Product.
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Genera un identificador único para el nuevo producto.
	p.ID = uuid.NewString()

	// Inserta el producto en la tabla products de PostgreSQL.
	_, erro := database.DB.Exec("INSERT INTO products (id, productname, stock, price) VALUES ($1, $2, $3, $4)", p.ID, p.ProductName, p.Stock, p.Price)

	if erro != nil {
		log.Printf("error: %v", erro)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "It was impossible to save product",
		})
		return
	}

	// Retorna el producto creado con estado 201 Created.
	c.JSON(http.StatusCreated, p)
}
