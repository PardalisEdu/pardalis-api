package personalization

import (
	"database/sql"
	"fmt"
	"gitlab.com/pardalis/pardalis-api/types"
)

// Store implementa PersonalizationStore
type Store struct {
	db *sql.DB
}

// NewStore crea una nueva instancia de Store
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetPersonalization obtiene la personalización de un usuario por su apodo
func (s *Store) GetPersonalization(apodo string) (*types.Personalization, error) {
	p := new(types.Personalization)
	err := s.db.QueryRow(
		"SELECT apodo, descripcion, foto, fecha_actualizacion FROM personalizacion WHERE apodo = ?",
		apodo,
	).Scan(&p.Apodo, &p.Descripcion, &p.Foto, &p.FechaActualizacion)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("personalization not found")
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

// CreatePersonalization crea una nueva personalización
func (s *Store) CreatePersonalization(p types.Personalization) error {
	_, err := s.db.Exec(
		"INSERT INTO personalizacion (apodo, descripcion, foto) VALUES (?, ?, ?)",
		p.Apodo, p.Descripcion, p.Foto,
	)
	return err
}

// UpdatePersonalization actualiza una personalización existente
func (s *Store) UpdatePersonalization(p types.Personalization) error {
	result, err := s.db.Exec(
		"UPDATE personalizacion SET descripcion = ?, foto = ? WHERE apodo = ?",
		p.Descripcion, p.Foto, p.Apodo,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("personalization not found")
	}

	return nil
}
