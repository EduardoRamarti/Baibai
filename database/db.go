package database

import (
	"database/sql"
	"fmt"
	"log"

	// Driver de PostgreSQL para Go usando pgx/v5.
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB es la conexión global a la base de datos PostgreSQL utilizada por toda la aplicación.
// Se inicializa en la función ConnectDB() al iniciar el servidor.
var DB *sql.DB

// ConnectDB establece una conexión con la base de datos PostgreSQL.
// Abre la conexión usando las credenciales y parámetros configurados en la cadena DSN.
// Verifica que la base de datos responda correctamente con un Ping.
// Si hay algún error, la aplicación se detiene inmediatamente.
func ConnectDB() {
	// DSN (Data Source Name) contiene:
	// - Usuario: mikoushi
	// - Contraseña: mikoushi
	// - Host: localhost
	// - Puerto: 5432 (puerto predeterminado de PostgreSQL)
	// - Base de datos: ecommercedb
	// - sslmode=disable: Sin encriptación SSL (solo para desarrollo)
	dsn := "postgres://mikoushi:mikoushi@localhost:5432/ecommercedb?sslmode=disable"

	var err error
	
	// Abre la conexión a PostgreSQL usando el driver pgx.
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	
	// Cierra la conexión cuando la función termina.
	defer DB.Close()

	// Verifica que la conexión a la base de datos sea exitosa enviando un ping.
	err = DB.Ping()
	if err != nil {
		log.Fatalf("El contenedor no responde: %v", err)
	}
	
	// Mensaje de confirmación cuando la conexión es exitosa.
	fmt.Println("¡Conexión a PostgreSQL (Alpine) exitosa!")

}
