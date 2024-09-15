// Package user 🐄 – Porque lo que más necesitamos en nuestras vidas es una tienda... de usuarios.
// Aquí es donde almacenamos esos valiosos datos, como contraseñas que deberían haber sido más fuertes. 💪
package user

import (
	"database/sql"
	"fmt"

	"codeberg.org/Pardalis/pardalis-api/types"
)

// Store 🐄 – El lugar donde mágicamente (o trágicamente) almacenamos nuestros usuarios.
// Esperemos que la base de datos no explote en el proceso... porque todos sabemos que las bases de datos NUNCA fallan, ¿cierto? 🙃
type Store struct {
	db *sql.DB // db 🐄 – Nuestra gloriosa base de datos. Porque claro, todo depende de ella... sin presión.
}

// NewStore 🐄 – Constructor básico que simplemente te entrega una Store y dice "aquí tienes, arreglatelas como puedas".
// ¿Esperabas algo más? Pff... sueña. 🌈
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// CreateUser 🐄 – Guarda un usuario en la base de datos con la ilusión de que todo saldrá bien.
// Porque insertar registros en SQL es siempre una operación de alto riesgo. 🎲
func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO usuarios (apodo, nombre, correo, contrasenna) VALUES (?, ?, ?, ?)", user.Apodo, user.Nombre, user.Correo, user.Contrasenna)
	if err != nil {
		return err // Si algo falla, no te preocupes, solo te devolveremos un error confuso. 🤷‍♂️
	}
	return nil // Si llegaste aquí, felicidades. Todo salió bien... probablemente. 🎉
}

// GetUserByCorreo 🐄 – Busca un usuario por su correo electrónico porque, obvio, eso nunca falla. ✉️
// Spoiler: A veces sí falla. Si el correo no existe, buena suerte con eso. 🤞
func (s *Store) GetUserByCorreo(correo string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM usuarios WHERE correo = ?", correo)
	if err != nil {
		return nil, err // Ups, algo salió mal... seguramente no es tu culpa. O sí. 🤔
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Apodo == "" { // Si el usuario no tiene apodo, básicamente no existe.
		return nil, fmt.Errorf("user not found") // ¡Sorpresa! El usuario no estaba ahí.
	}

	return u, nil
}

// GetUserByApodo 🐄 – Busca un usuario por su apodo. Porque todos los usuarios tienen apodos, ¿verdad? 🤷‍♀️
// Si no lo encuentras, es que probablemente no existe. Pero bueno, sigamos buscando.
func (s *Store) GetUserByApodo(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM usuarios WHERE apodo = ?", id)
	if err != nil {
		return nil, err // Si esto falla, solo te queda rezar. 🙏
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Apodo == "" { // El apodo es la clave... y si está vacío, adivina qué: no hay usuario.
		return nil, fmt.Errorf("user not found") // Aparentemente, este usuario no quiere ser encontrado. 🤫
	}

	return u, nil
}

// scanRowsIntoUser 🐄 – La función que toma filas de la base de datos y las convierte en un usuario.
// Porque los usuarios no pueden salir mágicamente de la base de datos. 🎩✨
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.Apodo,
		&user.Nombre,
		&user.Correo,
		&user.Contrasenna,
		&user.Registro,
	)
	if err != nil {
		return nil, err // Oh no, algo salió mal al convertir las filas en un usuario. 😱
	}

	return user, nil // Si todo salió bien, ¡felicidades! Has logrado obtener un usuario de la base de datos. 🎉
}
