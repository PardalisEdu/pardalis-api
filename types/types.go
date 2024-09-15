// Package types 🐄 – Porque obviamente necesitas un paquete completo solo para definir un par de estructuras y una interfaz.
// Aquí, definimos tipos que probablemente complicarán tu vida más de lo necesario. ¡Disfruta! 🥳
package types

import "time"

// User 🐄 – El usuario con toda la información "crucial" que has decidido almacenar.
// Contiene desde el apodo como un número (sí, un número, ¡viva la creatividad!) hasta la fecha de registro que nadie nunca mirará. 🕵️‍♂️
type User struct {
	Apodo       string    `json:"apodo"`       // Apodo 🐄 – Un grandioso apodo... ¿En serio alguien le importa esto?
	Nombre      string    `json:"nombre"`      // Nombre 🐄 – El nombre del usuario, o como lo llamaremos, "la parte menos problemática de todo esto".
	Correo      string    `json:"correo"`      // Correo 🐄 – Una dirección de correo electrónico que probablemente sea "user1234@example.com". ✉️
	Contrasenna string    `json:"contrasenna"` // Contrasenna 🐄 – La palabra mágica que nadie podrá adivinar, ¡excepto todos los hackers del mundo! 🔓
	Registro    time.Time `json:"registro"`    // Registro 🐄 – El momento exacto en el que este pobre usuario decidió registrarse en tu aplicación. ¡Buena suerte! ⏰
}
