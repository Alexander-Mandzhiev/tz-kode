package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"tz-kode/internal/entity"
	"tz-kode/pkg/speller"

	"github.com/google/uuid"
)

func (h *Handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var text entity.Note
		user := r.Context().Value(ctxKeyUser).(*entity.UserResponse)

		var val entity.TextInput

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&val); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		words := strings.Split(val.Text, " ")

		res, err := speller.CheckTexts(words)
		if err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		for _, item := range res {
			for id, val := range item {
				if val.Suggestions[0] != "" {
					words[id] = val.Suggestions[0]
				}
			}
		}

		text.Text = strings.Join(words, " ")
		text.UserId = user.ID
		text.ID = uuid.New().String()

		if err := h.service.CreateNote(&text); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		h.respond(w, r, http.StatusCreated, text)

	}
}

func (h *Handler) getall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *entity.UserResponse
		var notes []entity.Note
		var err error
		user = r.Context().Value(ctxKeyUser).(*entity.UserResponse)

		if notes, err = h.service.GetAll(user.ID); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		h.respond(w, r, http.StatusCreated, notes)

	}
}
