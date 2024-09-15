// Package types 🐄 – Porque obviamente necesitas un paquete completo solo para definir un par de estructuras y una interfaz.
// Aquí, definimos tipos que probablemente complicarán tu vida más de lo necesario. ¡Disfruta! 🥳
package types

// RegisterUserPayload 🐄 – La carga útil de registro que verifica que tus usuarios
// al menos tengan un nombre y correo, porque aparentemente eso es lo mínimo necesario para ser un ser humano. 😅
type RegisterUserPayload struct {
	Apodo       string `json:"apodo" validate:"required"`                     // Apodo 🐄 – Sí, ¿Realmente alguien escoge esto bien? ¿Los niños escogeran bien? O sera Daniel123. 🙃
	Nombre      string `json:"nombre" validate:"required"`                    // Nombre 🐄 – Obligatorio, porque nadie quiere un usuario sin nombre... ¿verdad?
	Correo      string `json:"correo" validate:"required,email"`              // Correo 🐄 – El correo del usuario, validado para asegurarse de que termine en "@", lo que podría ser suficiente. 🕵️‍♀️
	Contrasenna string `json:"contrasenna" validate:"required,min=3,max=130"` // Contrasenna 🐄 – La contraseña que va de 3 a 130 caracteres, porque todos sabemos que esos son los números mágicos para la seguridad. 🧙‍♂️
}

// LoginUserPayload 🐄 – La carga útil para iniciar sesión que define lo absolutamente
// mínimo necesario para que alguien pueda intentar acceder a tu aplicación.
// Porque, claro, solo necesitamos un correo y una contraseña para todo, ¿verdad? 🙄
type LoginUserPayload struct {
	Correo      string `json:"correo" validate:"required,email"` // Correo 🐄 – La dirección de correo electrónico que esperas que sea única, válida y recordada por todos los usuarios, porque es mucho pedir. 📧
	Contrasenna string `json:"contrasenna" validate:"required"`  // Contrasenna 🐄 – La contraseña que has decidido que debe ser requerida, pero por suerte para los hackers, no tienes reglas de complejidad. 🔑
}
