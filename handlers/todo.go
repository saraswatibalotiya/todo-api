package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"todo-api/database"
	"todo-api/models"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// CreateTodo creates a new TODO item
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = gocql.TimeUUID()
	todo.Created = time.Now()
	todo.Updated = time.Now()

	if err := database.Session.Query(`
        INSERT INTO items (id, user_id, title, description, status, created, updated)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
		todo.ID, todo.UserID, todo.Title, todo.Description, todo.Status, todo.Created, todo.Updated).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GetTodos retrieves a paginated list of TODO items
func GetTodos(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	status := r.URL.Query().Get("status")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset := (page - 1) * limit

	var todos []models.Todo
	var todo models.Todo

	query := "SELECT id, user_id, title, description, status, created, updated FROM items WHERE user_id = ?"
	params := []interface{}{userID}

	if status != "" {
		query += " AND status = ?"
		params = append(params, status)
	}
	query += " LIMIT ? OFFSET ?"
	params = append(params, limit, offset)

	iter := database.Session.Query(query, params...).Iter()
	for iter.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
		todos = append(todos, todo)
	}

	if err := iter.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GetTodo retrieves a TODO item by ID
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo

	if err := database.Session.Query(`
        SELECT id, user_id, title, description, status, created, updated FROM items WHERE id = ?`,
		params["id"]).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates a TODO item by ID
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID, _ = gocql.ParseUUID(params["id"])
	todo.Updated = time.Now()

	if err := database.Session.Query(`
        UPDATE items SET title = ?, description = ?, status = ?, updated = ?
        WHERE id = ?`, todo.Title, todo.Description, todo.Status, todo.Updated, todo.ID).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo deletes a TODO item by ID
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if err := database.Session.Query(`
        DELETE FROM items WHERE id = ?`, params["id"]).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
