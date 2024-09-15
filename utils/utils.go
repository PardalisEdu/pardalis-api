// Package utils 🐄 – La caja de herramientas de utilidades que todo programador necesita
// para que sus aplicaciones funcionen, o al menos para que no se rompan completamente.
// ¡Un verdadero espectáculo de soluciones "universales"! 🎩✨
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New() // Validate 🐄 – El validador que se encarga de asegurarse de que tus datos sean los correctos, o al menos que no sean completamente inútiles. 😅

// WriteJSON 🐄 – La función mágica que convierte tus datos en JSON y los envía al cliente.
// Porque todos sabemos que a nadie le gusta manejar errores directamente. 🙃
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json") // Porque la gente necesita saber que están recibiendo JSON, ¡como si no pudieran adivinarlo! 🙄
	w.WriteHeader(status)                              // Configura el código de estado HTTP, porque simplemente devolver 200 no siempre es suficiente. 🚀
	return json.NewEncoder(w).Encode(v)                // Codifica el valor en JSON y lo envía, porque nadie tiene tiempo para errores, ¡sólo para respuestas exitosas! 😅
}

// WriteError 🐄 – La función que envía un error como JSON. Porque hacer que los errores sean bonitos y presentables es la clave para la felicidad. 🌟
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()}) // Convierte el error en JSON, porque las cadenas de texto no tienen suficiente glamour. 💬
}

// ParseJSON 🐄 – La función que convierte el cuerpo de la solicitud en datos útiles.
// O, si te sientes aventurero, en datos completamente corruptos. ¡Es todo parte del viaje! 🎢
func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body") // Se queja si el cuerpo de la solicitud está vacío, porque eso siempre es divertido. 😆
	}

	return json.NewDecoder(r.Body).Decode(v) // Decodifica el cuerpo de la solicitud en el valor deseado, asumiendo que todo estará bien... ¡que el universo te acompañe! 🌌
}

// GetTokenFromRequest 🐄 – La función que saca el token de la solicitud. O, como llamamos a esto, "la búsqueda de la aguja en el pajar de tokens". 🧵
func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization") // Busca el token en los encabezados de autorización, porque eso tiene sentido, ¿verdad? 😏
	tokenQuery := r.URL.Query().Get("token")   // O busca el token en los parámetros de la URL, porque los desarrolladores aman tener múltiples lugares para almacenar secretos. 🤫

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return "" // No encontró el token, porque la vida es dura y a veces simplemente no hay recompensas. 🎲
}
