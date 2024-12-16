// Package auth 🐄 – La maravilla que maneja autenticaciones, tokens y esas cosas que
// todo el mundo necesita pero nadie realmente entiende. ¡Perfecto para confundir a tus usuarios y a ti mismo! 🎭
package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gitlab.com/pardalis/pardalis-api/configs"
	"gitlab.com/pardalis/pardalis-api/types"
	"gitlab.com/pardalis/pardalis-api/utils"
)

type contextKey string

const UserKey contextKey = "userApodo" // UserKey 🐄 – La llave mágica para encontrar a tu usuario en el contexto, porque todos necesitamos un poco de magia en nuestras vidas. 🪄

// WithJWTAuth 🐄 – El encantador middleware que intenta autenticar a los usuarios usando JWT.
// Porque nada dice "confianza" como agregar un token al encabezado y esperar lo mejor. 🕵️‍♂️
func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r) // Obtiene el token de la solicitud, si es que tienes uno... 🤷‍♂️

		token, err := validateJWT(tokenString) // Intenta validar el token, porque no hay nada más divertido que fallar en la validación. 😅
		if err != nil {
			log.Printf("failed to validate token: %v", err) // Registra el error, como si eso fuera a solucionar algo. 📜
			permissionDenied(w)                             // Niega el permiso con elegancia. 🛑
			return
		}

		if !token.Valid {
			log.Println("invalid token") // El token no es válido, ¡sorpresa! 🙄
			permissionDenied(w)          // Nuevamente, niega el permiso, porque no hay nada más que hacer. 🚪
			return
		}

		claims := token.Claims.(jwt.MapClaims)    // Extrae las reclamaciones del token, porque supongo que son importantes. 📜
		userApodo := claims["userApodo"].(string) // Obtiene el apodo del usuario, porque todos los usuarios tienen apodos... ¿no? 🤔

		u, err := store.GetUserByApodo(userApodo) // Intenta obtener al usuario por su apodo, como si esto fuera siempre exitoso. 😅
		if err != nil {
			log.Printf("failed to get user by id: %v", err) // Otro error registrado, como si eso fuera útil. 📜
			permissionDenied(w)                             // Y, por supuesto, negamos el permiso una vez más. 🚷
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.Apodo) // Añade el apodo del usuario al contexto, porque eso es lo que todos los programadores sueñan. 🌌
		r = r.WithContext(ctx)

		handlerFunc(w, r) // Llama a la función del manejador, porque eso es lo que se supone que debes hacer. 🎉
	}
}

// CreateJWT 🐄 – La función para crear tokens JWT, porque todos necesitamos más tokens en nuestras vidas.
// ¡Y este token probablemente durará más que tu última relación! 💔
func CreateJWT(secret []byte, userApodo string) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds) // Establece la expiración del token, porque nada dice "seguridad" como una fecha de vencimiento. 🗓️

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Crea un nuevo token, porque sí. 🎟️
		"userApodo": userApodo,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret) // Firma el token, porque un token sin firma es como un auto sin ruedas. 🚗
	if err != nil {
		return "", err // Retorna el error si algo sale mal, porque siempre hay algo que sale mal. 🤷‍♂️
	}

	return tokenString, err
}

// validateJWT 🐄 – La función que valida un token JWT, o como diría un desarrollador, la forma elegante de decir "hazlo funcionar". 🛠️
func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // Comprueba el método de firma, porque eso es lo más importante en la vida. 🔍
		}

		return []byte(configs.Envs.JWTSecret), nil // Devuelve la clave secreta, porque todos necesitamos un poco de misterio en nuestras vidas. 🎩
	})
}

// permissionDenied 🐄 – La función que maneja el caso en el que alguien no tiene permiso para hacer algo,
// o como diría tu terapeuta, "la forma más amable de decir que no". 🚫
func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied")) // Envía un error de permiso denegado, porque no hay nada más que hacer en este punto. 🙅‍♂️
}

// GetUserApodoFromContext 🐄 – Obtiene el apodo del usuario del contexto, porque claramente necesitas saber eso en algún momento. 🤷‍♀️
func GetUserApodoFromContext(ctx context.Context) string {
	if apodo, ok := ctx.Value(UserKey).(string); ok {
		return apodo
	}
	return ""
}

// VerifyJWT 🐄 – Esta función es el detective que revisa si el token JWT es válido o no. Si es válido,
// regresa los claims del token. Si no, regresa un error porque la autenticación ha fallado. 🔒
func VerifyJWT(tokenString string, secret []byte) (jwt.MapClaims, error) {
	// Definimos la función de verificación para obtener los claims del token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	// Si el token es válido y tiene los claims, los extraemos
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verificamos la expiración del token
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return nil, fmt.Errorf("token has expired")
			}
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
