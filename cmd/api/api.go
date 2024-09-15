// Package api 🐄 – Porque todos sabemos que necesitas otra API más en tu vida.
// ¡Sí, otra más para tu interminable lista de microservicios que nunca funcionan como deberían!
package api

// Porque sin una base de datos, ¿qué sería de nuestra vida?
import (
	"database/sql"
	"log"
	"net/http"

	"codeberg.org/Pardalis/pardalis-api/services/user"
	"github.com/gorilla/mux"
)

// APIServer 🐄 – El increíble servidor API que probablemente va a fallar
// tan pronto como intentes usarlo. Porque, sinceramente, ¿quién necesita más que
// una dirección y una conexión a la base de datos? ¡Claro, esto es suficiente! 🙄
type APIServer struct {
	addr string  // addr 🐄 – La dirección donde el servidor espera que los unicornios y las hadas de la red lo encuentren.
	db   *sql.DB // db 🐄 – La conexión a tu base de datos que mágicamente no debería tener problemas... nunca... 😅
}

// NewAPIServer 🐄 – El constructor más minimalista que jamás hayas visto.
// Porque, ¿para qué querrías lógica o configuración adicional? No, no, no.
// Aquí obtienes exactamente lo que ves: una dirección y una base de datos. 🎩✨
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Start 🏁 – Inicia el servidor API. O al menos eso intenta hacer... si todo sale bien. 🎉
// Configura las rutas, inicializa el almacenamiento de usuarios, y le pide a HTTP que escuche en un puerto.
// Si hay algún error, probablemente sea porque olvidaste levantar la base de datos... otra vez. 🤦‍♂️
func (s *APIServer) Start() error {
	// Creamos un nuevo enrutador que manejará todas las rutas. 🚗
	router := mux.NewRouter()

	// Creamos un subrouter específico para nuestra API versión 1. ¿Por qué? Bueno, porque "versionado" suena profesional. 📚
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Iniciamos la tienda de usuarios, que no tiene nada que ver con Amazon. 🛒
	userStore := user.NewStore(s.db)

	// Creamos el handler para los usuarios. Este será quien maneje todas esas solicitudes incómodas de registro. 🙇‍♂️
	userHandler := user.NewHandler(userStore)

	// Registramos todas las rutas relacionadas con usuarios, para que el subrouter pueda manejarlas como el ninja que es. 🥷
	userHandler.RegisterRoutes(subrouter)

	// El momento glorioso. Si llegamos hasta aquí sin explotar, el servidor está listo para atender las solicitudes. 🎉
	log.Printf("Servidor iniciado en el puerto %s\n", s.addr)

	// Ahora le decimos a HTTP que se ponga cómodo y escuche en la dirección y puerto que hemos configurado.
	// Si hay un error aquí, solo puedo desearte suerte. 🍀
	return http.ListenAndServeTLS(s.addr, "server.crt", "server.key", router)
}
