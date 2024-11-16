# Guía de Contribución - Pardalis Backend

¡Gracias por tu interés en contribuir a Pardalis Backend! Esta guía te ayudará a comprender nuestro proceso de desarrollo y las convenciones que seguimos.

## 📋 Índice
- [Configuración del Entorno](#configuración-del-entorno)
- [Convenciones de Código](#convenciones-de-código)
- [Proceso de Contribución](#proceso-de-contribución)
- [Estructura de Commits](#estructura-de-commits)
- [Pruebas](#pruebas)
- [Documentación](#documentación)

## Configuración del Entorno

### Requisitos Previos
- Go 1.23 o superior
- MySQL 8.0+
- golangci-lint (para linting)
- go-critic (para análisis de código)

### Configuración Inicial
1. Fork del repositorio
2. Clonar tu fork:
```bash
git clone https://github.com/tu-usuario/pardalis-api.git
cd pardalis-api
```

3. Añadir el repositorio upstream:
```bash
git remote add upstream https://codeberg.org/Pardalis/pardalis-api.git
```

4. Instalar dependencias:
```bash
go mod download
```

## Convenciones de Código

### Estilo de Código
Seguimos las convenciones estándar de Go:

```go
// Package user maneja la lógica de negocio relacionada con usuarios
package user

// User representa la entidad de usuario en el sistema
type User struct {
    ID        string    `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// GetByID obtiene un usuario por su ID
func GetByID(id string) (*User, error) {
    // Implementación
}
```

### Nombres de Archivos y Carpetas
- Usar snake_case para nombres de archivos: `user_repository.go`
- Usar nombres descriptivos y concisos
- Los archivos de prueba deben terminar en `_test.go`

## Proceso de Contribución

### 1. Planificación
- Revisa los issues existentes o crea uno nuevo
- Discute cambios mayores en los issues antes de comenzar

### 2. Desarrollo
1. Actualiza tu rama main:
```bash
git checkout main
git pull upstream main
```

2. Crea una rama para tu característica:
```bash
git checkout -b feature/nombre-descriptivo
```

3. Desarrolla siguiendo nuestras convenciones
4. Ejecuta las pruebas localmente
5. Ejecuta el linter:
```bash
golangci-lint run
```

### 3. Commit y Push
- Sigue nuestra estructura de commits
- Haz push a tu fork:
```bash
git push origin feature/nombre-descriptivo
```

## Estructura de Commits

Formato:
```
[Acción] Descripción concisa

Descripción detallada de los cambios (opcional)
```

Acciones permitidas:
- **[Añadido]**: Nueva funcionalidad
- **[Corregido]**: Corrección de errores
- **[Eliminado]**: Eliminación de código
- **[Característica]**: Característica completa
- **[Formateado]**: Cambios de formato
- **[Documentado]**: Cambios en documentación
- **[Optimizado]**: Mejoras de rendimiento
- **[Pruebas]**: Añadir o modificar pruebas

Ejemplos:
```
[Añadido] Implementación de rate limiter

- Añadido middleware de rate limiting
- Configuración por defecto de 100 req/min
- Tests unitarios para el rate limiter
```

```
[Corregido] Error en validación de JWT
```

## Pruebas

### Tipos de Pruebas
1. **Unitarias**: Para funciones y métodos individuales
2. **Integración**: Para interacciones entre componentes
3. **End-to-End**: Para flujos completos

### Escribiendo Pruebas
```go
func TestUserRepository_GetByID(t *testing.T) {
    tests := []struct {
        name    string
        id      string
        want    *User
        wantErr bool
    }{
        {
            name: "usuario existente",
            id:   "123",
            want: &User{ID: "123", Username: "test"},
        },
        {
            name:    "usuario no existente",
            id:      "456",
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Implementación
        })
    }
}
```

### Ejecutando Pruebas
```bash
# Todas las pruebas
go test ./...

# Pruebas con coverage
go test -cover ./...

# Pruebas específicas
go test ./pkg/user/...
```

## Documentación

### Comentarios de Código
- Cada paquete debe tener un comentario de documentación
- Las funciones exportadas deben estar documentadas
- Usa ejemplos en la documentación cuando sea útil

### Ejemplos en la Documentación
```go
// Example_handler_GetUser demuestra cómo usar el handler de usuario
func Example_handler_GetUser() {
    // Ejemplo de uso
}
```

### Actualizando la Documentación
- Mantén el README actualizado
- Documenta cambios en la API
- Actualiza la documentación de OpenAPI si es necesario

---

🍪 Has llegado al final de la guía de contribución. ¡Aquí tienes tu galleta virtual!