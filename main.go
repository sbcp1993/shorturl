package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shorturl/db"
	"shorturl/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var postgres *db.DBConnection
	var err error

	postgres, err = db.NewDBConnection()

	if err != nil {
		panic(err)
	} else if postgres == nil {
		panic("postgres is nil")
	}

	shandler := handlers.NewShorturl(postgres)

	sm := mux.NewRouter()
	sm.Handle("/generate", shandler).Methods("POST")
	sm.Handle("/{url}", shandler).Methods("GET")

	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	out := <-sigChan
	log.Println("graceful shutdown of service", out)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	cancel()
	s.Shutdown(ctx)
}
