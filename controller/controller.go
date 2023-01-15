package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
)

// ActorController struct represents the controller for actors
type ActorController struct {
	Model *ActorModel
}

// NewActorController creates a new ActorController
func NewActorController(db *sql.DB) *ActorController {
	return &ActorController{
		Model: &ActorModel{DB: db},
	}
}

// CreateActorHandler handles the creation of a new actor
func (c *ActorController) CreateActorHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var actor Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create actor
	id, err := c.Model.CreateActor(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with actor ID
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

// GetActorHandler handles the retrieval of an actor by ID
func (c *ActorController) GetActorHandler(w http.ResponseWriter, r *http.Request) {
	// Get actor ID from URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// Get actor
	actor, err := c.Model.GetActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with actor
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actor)
}

// ViewActorHandler handles the display of an actor's details
func ViewActorHandler(w http.ResponseWriter, r *http.Request) {
	// Get actor ID from URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// Get actor
	actor, err := ActorModel.GetActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse template
	t, err := template.ParseFiles("templates/actor.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute template
	err = t.Execute(w, actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

