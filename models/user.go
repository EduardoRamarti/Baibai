package models

// User representa la estructura de datos de un usuario en la aplicación.
// Se utiliza para almacenar y gestionar la información de los usuarios registrados.
type User struct {
	// ID es el identificador único del usuario (UUID).
	ID string `json:"id"`

	// FirstName es el nombre del usuario.
	FirstName string `json:"firstName"`

	// LastName es el apellido del usuario.
	LastName string `json:"lastName"`

	// Email es la dirección de correo electrónico del usuario (debe ser único).
	Email string `json:"email"`

	// Password es la contraseña del usuario. Se excluye de la serialización JSON (`json:"-"`).
	Password string `json:"-"`

	// UserName es el nombre de usuario (handle/username único) del usuario.
	UserName string `json:"username"`

	// UserType define el tipo o rol del usuario (ej: admin, customer, etc).
	UserType string `json:"usertype"`
}
