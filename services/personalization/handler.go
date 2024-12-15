package personalization

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/pardalis/pardalis-api/configs"
	"gitlab.com/pardalis/pardalis-api/services/auth"
	"gitlab.com/pardalis/pardalis-api/types"
	"gitlab.com/pardalis/pardalis-api/utils"
)

// Handler maneja las rutas relacionadas con la personalización
type Handler struct {
	store     types.PersonalizationStore
	userStore types.UserStore
}

// NewHandler crea una nueva instancia de Handler
func NewHandler(store types.PersonalizationStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

// RegisterRoutes registra las rutas del handler en el router
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{userApodo}/personalization",
		auth.WithJWTAuth(h.handleGetPersonalization, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/users/{userApodo}/personalization",
		auth.WithJWTAuth(h.handleUpdatePersonalization, h.userStore)).Methods(http.MethodPost, http.MethodPut)
}

// handleGetPersonalization maneja la obtención de la personalización de un usuario
func (h *Handler) handleGetPersonalization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userApodo := vars["userApodo"]

	// Verificar el token
	tokenString := utils.GetTokenFromRequest(r)
	if tokenString == "" {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing or invalid token"))
		return
	}

	claims, err := auth.VerifyJWT(tokenString, []byte(configs.Envs.JWTSecret))
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		return
	}

	// Verificar que el usuario solicita sus propios datos
	tokenApodo, ok := claims["userApodo"].(string)
	if !ok || tokenApodo != userApodo {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("unauthorized access"))
		return
	}

	// Obtener la personalización
	p, err := h.store.GetPersonalization(userApodo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := p.ToResponse()
	utils.WriteJSON(w, http.StatusOK, response)
}

// handleUpdatePersonalization maneja la actualización de la personalización de un usuario
func (h *Handler) handleUpdatePersonalization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userApodo := vars["userApodo"]

	// Verificar el token
	tokenString := utils.GetTokenFromRequest(r)
	if tokenString == "" {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing or invalid token"))
		return
	}

	claims, err := auth.VerifyJWT(tokenString, []byte(configs.Envs.JWTSecret))
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		return
	}

	// Verificar que el usuario modifica sus propios datos
	tokenApodo, ok := claims["userApodo"].(string)
	if !ok || tokenApodo != userApodo {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("unauthorized access"))
		return
	}

	// Decodificar el cuerpo de la petición
	var p types.Personalization
	if err := utils.ParseJSON(r, &p); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Asignar el apodo del usuario
	p.Apodo = userApodo

	// Verificar que el usuario existe
	_, err = h.userStore.GetUserByApodo(userApodo)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	// Intentar actualizar, si no existe crear uno nuevo
	err = h.store.UpdatePersonalization(p)
	if err != nil {
		if err.Error() == "personalization not found" {
			err = h.store.CreatePersonalization(p)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}
			utils.WriteJSON(w, http.StatusCreated, p.ToResponse())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, p.ToResponse())
}
