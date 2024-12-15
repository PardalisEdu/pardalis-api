package types

import "time"

// Personalization representa la estructura de datos para la personalización del usuario
type Personalization struct {
	Apodo              string    `json:"apodo"`
	Descripcion        string    `json:"descripcion"`
	Foto               string    `json:"foto"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

// PersonalizationStore define la interfaz para operaciones de la base de datos
type PersonalizationStore interface {
	GetPersonalization(apodo string) (*Personalization, error)
	CreatePersonalization(Personalization) error
	UpdatePersonalization(Personalization) error
}

// PersonalizationResponse es la estructura para las respuestas HTTP
type PersonalizationResponse struct {
	Descripcion string `json:"descripcion"`
	Foto        string `json:"foto"`
}

// ToResponse convierte un Personalization a PersonalizationResponse
func (p *Personalization) ToResponse() PersonalizationResponse {
	return PersonalizationResponse{
		Descripcion: p.Descripcion,
		Foto:        p.Foto,
	}
}
