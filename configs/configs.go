// Package configs 🐄 – Porque sin una buena configuración, ¿cómo vas a garantizar
// que tu aplicación se caiga en producción? Sí, otro archivo que manejarás mal.
// ¡Disfruta de tus variables de entorno mágicas! 🧙‍♂️✨
package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv" // godotenv 🐄 – La herramienta que asegura que tus secretos nunca se queden secretos.
)

// Config 🐄 – La configuración que asegura que tu aplicación funcione
// justo el tiempo suficiente para que lleguen las quejas de los usuarios.
// Contiene todos esos hermosos parámetros que se perderán en la nube. ☁️🔥
type Config struct {
	PublicHost             string // PublicHost 🐄 – Dónde estará "disponible" tu aplicación, asumiendo que a alguien le importe.
	Port                   string // Port 🐄 – El puerto favorito de tu aplicación, probablemente uno que ya esté en uso. 🙄
	DBUser                 string // DBUser 🐄 – Porque usar 'root' como usuario de la base de datos siempre es una buena idea, ¿verdad? 😅
	DBPassword             string // DBPassword 🐄 – La contraseña súper segura que compartes con todo el equipo... y con el mundo.
	DBAddress              string // DBAddress 🐄 – La dirección de tu base de datos. ¿Qué podría salir mal?
	DBName                 string // DBName 🐄 – El nombre de tu base de datos, porque 'default' es demasiado original.
	JWTSecret              string // JWTSecret 🐄 – ¡Shhh! Es un secreto... excepto cuando no lo es. 🤐
	JWTExpirationInSeconds int64  // JWTExpirationInSeconds 🐄 – Cuántos segundos durarán tus tokens JWT antes de expirar, o lo que es lo mismo, cuánto tiempo tienes hasta que todo se rompa. 🕒💥
}

// Envs 🐄 – Porque la palabra "environments" es demasiado larga.
// Contiene toda la magia negra que hace que tu aplicación intente funcionar.
var Envs = initConfig()

// initConfig 🐄 – Inicializa la configuración cargando tus variables de entorno
// desde un archivo .env, por si alguna vez te preguntas qué variables olvidaste
// configurar antes de subir a producción. ¡Sorpresa! 🎁
func initConfig() Config {
	godotenv.Load() // Carga el archivo .env, porque confiar en el entorno es sobrevalorado.

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),                                       // Configura el host público, que será ignorado por completo en producción.
		Port:                   getEnv("PORT", "8080"),                                                          // Selecciona un puerto... que probablemente ya esté en uso. 🎉
		DBUser:                 getEnv("DB_USER", "root"),                                                       // Usuario de la base de datos, porque 'root' es la elección de los campeones. 🏆
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),                                             // Contraseña ultra segura. Definitivamente nadie la adivinará. 🙄
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")), // Dirección de la base de datos, ¡esperemos que no haya cortafuegos! 🚧
		DBName:                 getEnv("DB_NAME", "padalis"),                                                    // El nombre de tu base de datos, ¿Por qué Pardalis tendra futuro? 🐄
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),                                // Un secreto tan seguro que lo estamos documentando aquí. 🤫
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),                             // Tiempo de expiración de los JWT, suficiente para que los hackers lo disfruten. 😈
	}
}

// getEnv 🐄 – Obtiene el valor de la variable de entorno o se rinde y usa el valor por defecto.
// Porque, seamos sinceros, nunca recordarás configurarlas todas. 🤷‍♂️
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// getEnvAsInt 🐄 – Obtiene la variable de entorno como un número entero o algo que se le parezca.
// No garantiza que el valor tenga sentido, pero, ¿quién necesita sentido en este punto? 🤡
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
