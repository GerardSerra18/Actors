package controller

import (
	"Actors/model"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
)

// ActorController struct represents the controller for actors
type ActorController struct {
	Model model.ActorModel
}

// NewActorController creates a new ActorController
func NewActorController(db *sql.DB) *ActorController {
	return &ActorController{
		Model: model.ActorModel{DB: db},
	}
}

// CreateActorHandler handles the creation of a new actor
func (c *ActorController) CreateActorHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var actor model.Actor
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
func (c *ActorController) ViewActorHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	actor, _ := c.Model.GetActor(id)
	t, _ := template.ParseFiles("views/actor.html")
	t.Execute(w, actor)
}

// UpdateActorHandler handles the update of an actor's details
func (c *ActorController) UpdateActorHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var actor model.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update actor
	err = c.Model.UpdateActor(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Write([]byte("Actor updated successfully"))
}

// DeleteActorHandler handles the deletion of an actor by ID
func (c *ActorController) DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
	// Get actor ID from URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// Delete actor
	err = c.Model.DeleteActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Write([]byte("Actor deleted successfully"))
}
