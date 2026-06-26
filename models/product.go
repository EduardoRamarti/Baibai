package models

// Product representa la estructura de datos de un producto en la aplicación.
// Se utiliza para serializar/deserializar datos entre JSON y la base de datos.
type Product struct {
	// ID es el identificador único del producto (UUID).
	ID string `json:"id"`

	// ProductName es el nombre o descripción del producto.
	ProductName string `json:"productName"`

	// Stock es la cantidad de unidades disponibles del producto.
	Stock int32 `json:"stock"`

	// Price es el precio del producto (en centavos o unidad mínima).
	Price int64 `json:"price"`
}
