package user

import (
	"fmt"
	"net/http"

	"codeberg.org/Pardalis/pardalis-api/configs"
	"codeberg.org/Pardalis/pardalis-api/services/auth"
	"codeberg.org/Pardalis/pardalis-api/types"
	"codeberg.org/Pardalis/pardalis-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Handler 🐄 – El valiente guardián de nuestras rutas de usuario. Está aquí para manejar las solicitudes
// de registro, inicio de sesión y obtención de usuario. Sí, porque solo él puede salvarnos de la confusión. 🌟
type Handler struct {
	store types.UserStore
}

// NewHandler 🐄 – El creador de nuestro héroe manejador. Al parecer, hay alguien que necesita ser responsable
// de las solicitudes de usuario, y este es el elegido. 🏆
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes 🐄 – El gran registrador de rutas. Aquí es donde se configuran las rutas para el manejo de
// usuarios, porque el enrutamiento es una tarea que alguien tiene que hacer. 🚦
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/users/{userApodo}", auth.WithJWTAuth(h.handleGetUser, h.store)).Methods(http.MethodGet)
}

// handleLogin 🐄 – El mago del inicio de sesión. Aquí es donde intentamos iniciar sesión, verificar
// credenciales y, si todo va bien, devolver un token. Si algo sale mal, el mundo se desmorona. 😱
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.store.GetUserByCorreo(user.Correo)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Contrasenna, []byte(user.Contrasenna)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.Apodo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
	if err != nil {
		return
	}
}

// handleRegister 🐄 – El héroe del registro. Aquí registramos a un nuevo usuario, comprobamos si ya existe,
// y si no, procedemos a guardar la nueva creación en la base de datos. Porque ¿qué sería de nosotros sin
// nuevos usuarios? 🎉
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetUserByCorreo(user.Correo)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Correo))
		return
	}

	hashedPassword, err := auth.HashPassword(user.Contrasenna)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		Apodo:       user.Apodo,
		Nombre:      user.Nombre,
		Correo:      user.Correo,
		Contrasenna: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, nil)
	if err != nil {
		return
	}
}

// handleGetUser 🐄 – El guardián de la información del usuario. Aquí obtenemos los detalles de un usuario
// usando su apodo. Porque todos necesitamos saber quién es quién, ¿verdad? 👀
func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userApodo, ok := vars["userApodo"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	tokenString := utils.GetTokenFromRequest(r)
	if tokenString == "" {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing or invalid token"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	claims, err := auth.VerifyJWT(tokenString, secret)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error verifying token: %v", err))
		return
	}

	tokenApodo, ok := claims["userApodo"].(string)
	if !ok || tokenApodo != userApodo {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("you are not authorized to view this user's information"))
		return
	}

	user, err := h.store.GetUserByApodo(userApodo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Convertir a la estructura de respuesta antes de enviar
	response := user.ToResponse()
	err = utils.WriteJSON(w, http.StatusOK, response)
	if err != nil {
		return
	}
}
