package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "UP")
}

func TestPingHandler(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestApiStatus(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/status", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "operational")
}

func TestEchoHandler(t *testing.T) {
	router := setupRouter()
	
	tests := []struct {
		name     string
		payload  string
		expected string
	}{
		{"Normal", , "hello"},
		{"Empty", , ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/echo", strings.NewReader(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)
			
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Contains(t, w.Body.String(), tt.expected)
			assert.Contains(t, w.Body.String(), "reverse")
		})
	}
}
