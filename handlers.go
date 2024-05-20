package api

import (
	"encoding/json"
	"net/http"
	"time"

	"example.com/todo-app-go/internal/auth"
	"example.com/todo-app-go/internal/mock"
	"example.com/todo-app-go/pkg/models"
)

// APIHandler handles API requests
type APIHandler struct {
	mockService *mock.MockService
	authService *auth.AuthService
}

// NewAPIHandler creates a new instance of APIHandler
func NewAPIHandler(mockService *mock.MockService, authService *auth.AuthService) *APIHandler {
	return &APIHandler{
		mockService: mockService,
		authService: authService,
	}
}

// Login handles user login
func (handler *APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := handler.authService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// CreateToDoList handles creation of a new TO-DO list
func (handler *APIHandler) CreateToDoList(w http.ResponseWriter, r *http.Request) {
	var newList models.ToDoList
	if err := json.NewDecoder(r.Body).Decode(&newList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newList.CreationDate = time.Now()
	newList.ModificationDate = newList.CreationDate

	handler.mockService.AddToDoList(newList)

	w.WriteHeader(http.StatusCreated)
}

// DeleteToDoList handles deletion of a T
func (handler *APIHandler) DeleteToDoList(w http.ResponseWriter, r *http.Request) {
	var listID int
	if err := json.NewDecoder(r.Body).Decode(&listID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement deletion logic based on your requirements

	w.WriteHeader(http.StatusOK)
}

// AddToDoMessage handles addition of a new TO-DO message
func (handler *APIHandler) AddToDoMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage models.ToDoMessage
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMessage.CreationDate = time.Now()
	newMessage.ModificationDate = newMessage.CreationDate

	handler.mockService.AddToDoMessage(newMessage)

	w.WriteHeader(http.StatusCreated)
}

// DeleteToDoMessage handles deletion of a TO-DO message
func (handler *APIHandler) DeleteToDoMessage(w http.ResponseWriter, r *http.Request) {
	var messageID int
	if err := json.NewDecoder(r.Body).Decode(&messageID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement deletion logic based on your requirements

	w.WriteHeader(http.StatusOK)
}

// UpdateToDoMessage handles update of a TO-DO message
func (handler *APIHandler) UpdateToDoMessage(w http.ResponseWriter, r *http.Request) {
	var updatedMessage models.ToDoMessage
	if err := json.NewDecoder(r.Body).Decode(&updatedMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMessage.ModificationDate = time.Now()

	// Implement update logic based on your requirements

	w.WriteHeader(http.StatusOK)
}

// GetToDoLists handles retrieval of all TO-DO lists
func (handler *APIHandler) GetToDoLists(w http.ResponseWriter, r *http.Request) {
	allLists := handler.mockService.GetAllToDoLists()

	responseData, err := json.Marshal(allLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
