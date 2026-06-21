package models

type Product struct {
	ID          string `json:"id"`
	ProductName string `json:"productName"`
	Stock       int32  `json:"stock"`
	Price       int64  `json:"price"`
}
