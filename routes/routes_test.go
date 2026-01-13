package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/helioLJ/go-booking-rest-api/db"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() {
	// Use a separate test database
	os.Remove("test.db")
	var err error
	db.DB, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to open test database")
	}

	db.DB.SetMaxOpenConns(10)
	db.DB.SetMaxIdleConns(5)

	// Create tables
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	db.DB.Exec(createUsersTable)

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	db.DB.Exec(createEventsTable)

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	db.DB.Exec(createRegistrationsTable)
}

func teardownTestDB() {
	db.DB.Close()
	os.Remove("test.db")
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	RegisterRoutes(router)
	return router
}

func TestSignupEndpoint(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := setupRouter()

	// Test successful signup
	signupData := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonData, _ := json.Marshal(signupData)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "User created successfully!" {
		t.Errorf("Unexpected response message: %v", response["message"])
	}
}

func TestLoginEndpoint(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := setupRouter()

	// First, create a user
	signupData := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonData, _ := json.Marshal(signupData)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Now test login
	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["token"] == nil {
		t.Error("Expected token in response")
	}
}

func TestLoginWithWrongPassword(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := setupRouter()

	// First, create a user
	signupData := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonData, _ := json.Marshal(signupData)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Try login with wrong password
	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "wrongpassword",
	}
	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestGetEventsEndpoint(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := setupRouter()

	req, _ := http.NewRequest("GET", "/events", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Body should contain valid JSON (array)
	if w.Body.Len() == 0 {
		t.Error("Expected non-empty response body")
	}
}

func TestCreateEventWithoutAuth(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := setupRouter()

	eventData := map[string]interface{}{
		"name":        "Test Event",
		"description": "Test Description",
		"location":    "Test Location",
		"dateTime":    "2024-09-15T09:00:00Z",
	}
	jsonData, _ := json.Marshal(eventData)

	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Should fail without authentication
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}
