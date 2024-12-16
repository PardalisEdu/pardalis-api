// utils/slug.go
package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(title string) string {
	// Convertir a minúsculas
	slug := strings.ToLower(title)

	// Reemplazar caracteres especiales y espacios con guiones
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Eliminar guiones del inicio y final
	slug = strings.Trim(slug, "-")

	return slug
}
