package main

import (
	//"database/sql"
	"log"
	"net/http"
	//"os"
	"time"
	"Actors/controller"
	_"github.com/lib/pq"
)

func main() {
    // Connect to database
    db, err := connectToDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize database
    if err = initDB(); err != nil {
        log.Fatal(err)
    }
    
    // Create actor controller
    actorCtrl := controller.NewActorController(db)

    // Create router
    mux := http.NewServeMux()
    mux.HandleFunc("/actors/create", actorCtrl.CreateActorHandler)
	mux.HandleFunc("/actors/get", actorCtrl.GetActorHandler)
    mux.HandleFunc("/actors/view", actorCtrl.ViewActorHandler)
    mux.HandleFunc("/actors/update", actorCtrl.UpdateActorHandler)
    mux.HandleFunc("/actors/delete", actorCtrl.DeleteActorHandler)
    mux.Handle("/actors/images/", http.StripPrefix("/actors/images/", http.FileServer(http.Dir("actors/images"))))

    // Create server
    server := &http.Server{
        Addr:         ":8000",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    // Start server
    log.Println("Starting server on :8000")
    log.Fatal(server.ListenAndServe())
}
