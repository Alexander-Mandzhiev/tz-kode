package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"tz-kode/internal/handlers"
	"tz-kode/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestHandlers_AuthenticatedUser(t *testing.T) {
	u := TestUser(t)
	repo, teardown := TestRepository(t, databaseURL)
	defer teardown("users")
	services := service.NewService(repo)
	repo.User.Create(u)

	testCases := []struct {
		name         string
		coodieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name:         "authenticated",
			coodieValue:  map[interface{}]interface{}{"user_id": u.ID},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			coodieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	handler := handlers.NewHandler(services, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	hand := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/", nil)

			cookieStr, _ := sc.Encode("session", tc.coodieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", "session", cookieStr))
			handler.AuthenticateUser(hand).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})

	}
}

func TestHandlers_SignUp(t *testing.T) {

	repo, teardown := TestRepository(t, databaseURL)
	services := service.NewService(repo)

	defer teardown("users")
	handler := handlers.NewHandler(services, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "Password@123",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid paramd",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signup", b)
			handler.InitRouters().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestHandlers_SignIn(t *testing.T) {

	u := TestUser(t)
	repo, teardown := TestRepository(t, databaseURL)
	services := service.NewService(repo)

	repo.User.Create(u)

	defer teardown("users")
	handler := handlers.NewHandler(services, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signin", b)
			handler.InitRouters().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})

	}
}

func TestHandlers_CreateNote(t *testing.T) {

	u := TestUser(t)
	repo, teardown := TestRepository(t, databaseURL)
	services := service.NewService(repo)

	repo.User.Create(u)

	defer teardown("users")
	handler := handlers.NewHandler(services, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "unauthorize",
			payload: map[string]string{
				"id":         uuid.New().String(),
				"text":       "text example",
				"user_id":    u.ID,
				"created_at": time.Now().GoString(),
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signin", b)
			handler.InitRouters().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
