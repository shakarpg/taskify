package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"taskify/internal/models"
	"taskify/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	r := router.NewRouter()
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestLoginEndpoint(t *testing.T) {
	r := router.NewRouter()

	body := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)
	assert.NotEmpty(t, response["token"])
}

func TestCreateTaskWithoutAuth(t *testing.T) {
	r := router.NewRouter()

	task := models.Task{Title: "Test Task"}
	jsonBody, _ := json.Marshal(task)

	req := httptest.NewRequest("POST", "/api/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Deve retornar 401 sem autenticação
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
