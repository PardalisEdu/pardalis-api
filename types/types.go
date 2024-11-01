// Package types 🐄 – Porque obviamente necesitas un paquete completo solo para definir un par de estructuras y una interfaz.
// Aquí, definimos tipos que probablemente complicarán tu vida más de lo necesario. ¡Disfruta! 🥳
package types

import "time"

// User 🐄 – El usuario con toda la información "crucial" que has decidido almacenar.
// Contiene desde el apodo como un número (sí, un número, ¡viva la creatividad!) hasta la fecha de registro que nadie nunca mirará. 🕵️‍♂️
type User struct {
	Apodo       string    `json:"apodo"`
	Nombre      string    `json:"nombre"`
	Correo      string    `json:"correo"`
	Contrasenna string    `json:"-"`
	Registro    time.Time `json:"-"`
}

// UserResponse - Estructura específica para respuestas HTTP
type UserResponse struct {
	Apodo  string `json:"apodo"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

// ToResponse - Convierte un User a UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		Apodo:  u.Apodo,
		Nombre: u.Nombre,
		Correo: u.Correo,
	}
}
