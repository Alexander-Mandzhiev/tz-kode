package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"tz-kode/internal/entity"

	"github.com/google/uuid"
)

const (
	sessionName        = "session"
	ctxKeyUser  ctxKey = iota
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int8

type UserExist struct {
	Message string `json:"message"`
}

func (h *Handler) signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &entity.UserRequest{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		if user, _ := h.service.FindByEmain(req.Email); user != nil {
			exist := &UserExist{
				Message: "user with this email already exists",
			}
			h.respond(w, r, http.StatusBadRequest, exist)
			return
		}

		u := &entity.User{
			ID:       uuid.New().String(),
			Email:    req.Email,
			Password: req.Password,
			Username: "",
		}
		if err := h.service.Create(u); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		session, err := h.sessionStore.Get(r, sessionName)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID

		if err := h.sessionStore.Save(r, w, session); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		u.ClearPassword()
		h.respond(w, r, http.StatusCreated, nil)
	}
}

func (h *Handler) signin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &entity.UserRequest{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := h.service.FindByEmain(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			h.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := h.sessionStore.Get(r, sessionName)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID

		if err := h.sessionStore.Save(r, w, session); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, nil)
	}
}

func (h *Handler) AuthenticateUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := h.sessionStore.Get(r, sessionName)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			h.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := h.service.FindById(id.(string))
		if err != nil {
			h.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		req := &entity.UserResponse{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, req)))
	})
}

func (h *Handler) whoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*entity.UserResponse))
	}
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
