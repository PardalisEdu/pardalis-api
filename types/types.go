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

type Blog struct {
	ID               string    `json:"id"`
	Titulo           string    `json:"titulo"`
	Slug             string    `json:"slug"`
	Contenido        string    `json:"contenido"`
	Extracto         string    `json:"extracto"`
	ImagenPortada    string    `json:"imagen_portada"`
	FechaPublicacion time.Time `json:"fecha_publicacion"`
	Estado           string    `json:"estado"`
	Categoria        string    `json:"categoria"`
	TiempoLectura    int       `json:"tiempo_lectura"`
	AutorApodo       string    `json:"autor_apodo"`
	MetaDescripcion  string    `json:"meta_descripcion"`
	MetaKeywords     string    `json:"meta_keywords"`
	Tags             []string  `json:"tags"`
}
