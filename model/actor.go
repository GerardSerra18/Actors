package model

import (
	"database/sql"
)

// Actor struct represents an actor
type Actor struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	AudienceRating int `json:"audience_rating"`
}



// ActorModel struct represents the Actor model
type ActorModel struct {
	DB *sql.DB
}

// CreateActor creates a new actor in the database
func (a *ActorModel) CreateActor(actor *Actor) (int, error) {
	// Insert query and parameters
	query := "INSERT INTO actors (first_name, last_name, gender, age) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := a.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute query
	err = stmt.QueryRow(actor.FirstName, actor.LastName, actor.Gender, actor.Age).Scan(&actor.ID)
	if err != nil {
		return 0, err
	}
	return actor.ID, nil
}

// GetActor retrieves an actor from the database by ID
func (a *ActorModel) GetActor(id int) (*Actor, error) {
	// Select query and parameters
	query := "SELECT id, first_name, last_name, gender, age FROM actors WHERE id = $1"
	stmt, err := a.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	actor := &Actor{}
	err = stmt.QueryRow(id).Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.Gender, &actor.Age)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

// UpdateActor updates an actor in the database
func (a *ActorModel) UpdateActor(actor *Actor) error {
    // Update query and parameters
    query := "UPDATE actors SET first_name = $1, last_name = $2, gender = $3, age = $4 WHERE id = $5"
    stmt, err := a.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute query
    _, err = stmt.Exec(actor.FirstName, actor.LastName, actor.Gender, actor.Age, actor.ID)
    if err != nil {
        return err
    }
    return nil
}


// DeleteActor deletes an actor from the database by ID
func (a *ActorModel) DeleteActor(id int) error {
    // Delete query and parameters
    query := "DELETE FROM actors WHERE id = $1"
    stmt, err := a.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute query
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
