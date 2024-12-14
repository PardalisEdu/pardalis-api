package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gitlab.com/pardalis/pardalis-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

// NewMySQLStorage 🐄 – Porque necesitamos una base de datos más robusta que SQLite.
// ¡Bienvenido a MySQL, donde las conexiones son más complejas pero al menos es "enterprise"! 🏢
func NewMySQLStorage() (*sql.DB, error) {
	// Construimos el DSN (Data Source Name) con los datos de configuración
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		configs.Envs.DBUser,
		configs.Envs.DBPassword,
		configs.Envs.DBAddress,
		configs.Envs.DBName,
	)

	// Intentamos abrir la conexión a MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error connecting to MySQL: %v", err)
		return nil, err
	}

	// Configuramos el pool de conexiones
	db.SetMaxOpenConns(25)                 // Máximo de conexiones abiertas
	db.SetMaxIdleConns(25)                 // Máximo de conexiones inactivas
	db.SetConnMaxLifetime(5 * time.Minute) // Tiempo máximo de vida de una conexión

	// Verificamos la conexión
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging MySQL: %v", err)
		return nil, err
	}

	return db, nil
}

// InitializeDatabase 🐄 – Función para crear las tablas necesarias si no existen
func InitializeDatabase(db *sql.DB) error {
	// SQL para crear la tabla de usuarios
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS usuarios (
		apodo VARCHAR(255) PRIMARY KEY,
		nombre VARCHAR(255) NOT NULL,
		correo VARCHAR(255) UNIQUE NOT NULL,
		contrasenna VARCHAR(255) NOT NULL,
		registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// Ejecutamos el SQL
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		return err
	}

	log.Println("Database tables initialized successfully")
	return nil
}
