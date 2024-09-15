package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// NewSQLiteStorage 🐄 – Porque, ¿quién necesita una base de datos robusta cuando puedes usar SQLite?
// ¡Es ligera, portátil, y perfecta para evitar problemas... hasta que no lo sea! 🙃
// dbPath: La ruta al archivo de la base de datos SQLite, porque todos sabemos que te encanta jugar con rutas de archivos.
// Devuelve: *sql.DB (la conexión a tu nueva base de datos, que esperamos no se corrompa de inmediato), o un error si algo explota, lo cual es altamente probable.
func NewSQLiteStorage(dbPath string) (*sql.DB, error) {
	// Intentamos abrir la base de datos como si todo fuera a salir bien. 🕊️
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		// Si algo sale mal, hacemos lo más lógico: quejarse en el log y terminar el programa. 🤷‍♂️
		log.Fatal(err)
	}

	// ¡Mira, has abierto la base de datos! Esperemos que dure. 🎉
	return db, nil
}
