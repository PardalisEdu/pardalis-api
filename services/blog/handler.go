package blog

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/pardalis/pardalis-api/services/auth"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/pardalis/pardalis-api/types"
	"gitlab.com/pardalis/pardalis-api/utils"
)

type Handler struct {
	store     types.BlogStore
	userStore types.UserStore
}

func NewBlogHandler(store types.BlogStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/blogs", h.handleGetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{slug}", h.handleGetBlog).Methods("GET")
	router.HandleFunc("/blogs", auth.WithJWTAuth(h.handleCreateBlog, h.userStore)).Methods("POST")
	router.HandleFunc("/blogs/{id}", auth.WithJWTAuth(h.handleUpdateBlog, h.userStore)).Methods("PUT")
	router.HandleFunc("/blogs/{id}", auth.WithJWTAuth(h.handleDeleteBlog, h.userStore)).Methods("DELETE")
}

func (h *Handler) handleGetBlogs(w http.ResponseWriter, r *http.Request) {
	// Obtener parámetros de paginación y filtrado
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 50 {
		limit = 10
	}

	categoria := r.URL.Query().Get("categoria")

	blogs, err := h.store.GetBlogs(page, limit, categoria)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, blogs)
	if err != nil {
		return
	}
}

func (h *Handler) handleGetBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	blog, err := h.store.GetBlogBySlug(slug)
	if err != nil {
		if err.Error() == "blog not found" {
			utils.WriteError(w, http.StatusNotFound, err)
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, blog)
	if err != nil {
		return
	}
}

func (h *Handler) handleCreateBlog(w http.ResponseWriter, r *http.Request) {

	// Obtener el usuario del contexto (añadido por el middleware de autenticación)
	autorApodo := auth.GetUserApodoFromContext(r.Context())
	if autorApodo == "" { // Ahora verificamos si está vacío
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	// Parsear y validar el payload
	var payload types.CreateBlogPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Crear el slug desde el título
	slug := utils.GenerateSlug(payload.Titulo)

	// Crear el blog
	blog := types.Blog{
		ID:               uuid.New().String(),
		Titulo:           payload.Titulo,
		Slug:             slug,
		Contenido:        payload.Contenido,
		Extracto:         payload.Extracto,
		ImagenPortada:    payload.ImagenPortada,
		FechaPublicacion: time.Now(),
		Estado:           "borrador", // Por defecto es borrador
		Categoria:        payload.Categoria,
		TiempoLectura:    payload.TiempoLectura,
		AutorApodo:       autorApodo,
		MetaDescripcion:  payload.MetaDescripcion,
		MetaKeywords:     payload.MetaKeywords,
		Tags:             payload.Tags,
	}

	println("PASO 4")
	// Guardar en la base de datos
	err := h.store.CreateBlog(blog)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	println("PASO 5")
	err = utils.WriteJSON(w, http.StatusCreated, blog)
	if err != nil {
		return
	}

	println("PASO 6")
}

func (h *Handler) handleUpdateBlog(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del blog
	vars := mux.Vars(r)
	blogID := vars["id"]

	// Obtener el usuario del contexto
	autorApodo := auth.GetUserApodoFromContext(r.Context())
	if autorApodo == "" { // Ahora verificamos si está vacío
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	// Obtener el blog actual para verificar permisos
	currentBlog, err := h.store.GetBlogByID(blogID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("blog not found"))
		return
	}

	// Verificar que el usuario es el autor del blog
	if currentBlog.AutorApodo != autorApodo {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("not authorized to update this blog"))
		return
	}

	// Parsear y validar el payload
	var payload types.UpdateBlogPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Actualizar solo los campos proporcionados
	if payload.Titulo != "" {
		currentBlog.Titulo = payload.Titulo
		currentBlog.Slug = utils.GenerateSlug(payload.Titulo)
	}
	if payload.Contenido != "" {
		currentBlog.Contenido = payload.Contenido
	}
	if payload.Extracto != "" {
		currentBlog.Extracto = payload.Extracto
	}
	if payload.ImagenPortada != "" {
		currentBlog.ImagenPortada = payload.ImagenPortada
	}
	if payload.Categoria != "" {
		currentBlog.Categoria = payload.Categoria
	}
	if payload.TiempoLectura != 0 {
		currentBlog.TiempoLectura = payload.TiempoLectura
	}
	if payload.Estado != "" {
		currentBlog.Estado = payload.Estado
	}
	if payload.MetaDescripcion != "" {
		currentBlog.MetaDescripcion = payload.MetaDescripcion
	}
	if payload.MetaKeywords != "" {
		currentBlog.MetaKeywords = payload.MetaKeywords
	}
	if len(payload.Tags) > 0 {
		currentBlog.Tags = payload.Tags
	}

	// Guardar los cambios
	err = h.store.UpdateBlog(*currentBlog)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, currentBlog)
	if err != nil {
		return
	}
}

func (h *Handler) handleDeleteBlog(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del blog
	vars := mux.Vars(r)
	blogID := vars["id"]

	// Obtener el usuario del contexto
	autorApodo := auth.GetUserApodoFromContext(r.Context())
	if autorApodo == "" { // Ahora verificamos si está vacío
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	// Verificar que el blog existe y el usuario es el autor
	currentBlog, err := h.store.GetBlogByID(blogID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("blog not found"))
		return
	}

	if currentBlog.AutorApodo != autorApodo {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("not authorized to delete this blog"))
		return
	}

	// Eliminar el blog
	err = h.store.DeleteBlog(blogID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Blog deleted successfully"})
	if err != nil {
		return
	}
}
