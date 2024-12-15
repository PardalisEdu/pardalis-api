package middleware

import (
	"net/http"
	"strings"

	"github.com/rs/cors"
)

// CorsConfig contiene la configuración para CORS
type CorsConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	MaxAge           int
	AllowCredentials bool
}

// DefaultCorsConfig retorna una configuración CORS por defecto
func DefaultCorsConfig() *CorsConfig {
	return &CorsConfig{
		AllowedOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		MaxAge:           300,
		AllowCredentials: true,
	}
}

// CORS middleware maneja las cabeceras CORS
func CORS(config *CorsConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Validar origen
			allowedOrigin := ""
			for _, o := range config.AllowedOrigins {
				if o == "*" || o == origin {
					allowedOrigin = o
					break
				}
			}

			if allowedOrigin == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Establecer cabeceras CORS
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Manejar preflight requests
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Methods",
					strings.Join(config.AllowedMethods, ","))
				w.Header().Set("Access-Control-Allow-Headers",
					strings.Join(config.AllowedHeaders, ","))
				w.Header().Set("Access-Control-Max-Age",
					string(config.MaxAge))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func NewCorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://localhost:4173"}, // Orígenes permitidos
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
			"X-Requested-With",
		},
		AllowCredentials: true,
		Debug:            false, // Habilita logs de depuración
	})
}
