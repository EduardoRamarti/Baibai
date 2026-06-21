package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDB() {
	// postgres://usuario:password@host:puerto/basededatos?sslmode=disable
	dsn := "postgres://mikoushi:mikoushi@localhost:5432/ecommercedb?sslmode=disable"

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatalf("El contenedor no responde: %v", err)
	}
	fmt.Println("¡Conexión a PostgreSQL (Alpine) exitosa!")

}
