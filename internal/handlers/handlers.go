package handlers

import (
	"tz-kode/internal/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Handler struct {
	service      *service.Service
	sessionStore sessions.Store
}

func NewHandler(service *service.Service, sessionStore sessions.Store) *Handler {
	return &Handler{
		service:      service,
		sessionStore: sessionStore,
	}
}

func (h *Handler) InitRouters() *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	router.HandleFunc("/signup", h.signup()).Methods("POST")
	router.HandleFunc("/signin", h.signin()).Methods("POST")

	private := router.PathPrefix("/profile").Subrouter()
	private.Use(h.AuthenticateUser)
	private.HandleFunc("/whoami", h.whoami()).Methods("GET")
	private.HandleFunc("/notes", h.create()).Methods("POST")
	private.HandleFunc("/notes", h.getall()).Methods("GET")

	return router
}
