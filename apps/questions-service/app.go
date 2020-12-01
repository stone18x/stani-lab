package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"com.stani.questionsservice/services"
	"github.com/gorilla/mux"
)

// App contains the webserver
type App struct {
	Router *mux.Router
}

// Run starts the webserver
func (a *App) Run(port int) {
	a.Router.StrictSlash(true)
	subRouter := a.Router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/questions", returnQuestions).Methods("GET")
	log.Printf("Go web server started at port: %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

func returnQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Invoked /questions")
	w.Header().Set("Content-Type", "application/json")
	questionsService := services.NewQuestionsService()
	json.NewEncoder(w).Encode(questionsService.Get())
}
