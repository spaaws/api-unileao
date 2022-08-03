package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []Event

var events = allEvents{
	{ID: 1, Title: "Unileão DAY", Description: "Evento para alunos das escolas"},
	{ID: 2, Title: "App Projeto Integrador", Description: "Apresentação do projeto integrador"},
}

func main() {
	//access log
	log.Println("Starting API...")
	//endpoints
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/events", GetAllEvents).Methods("GET")
	// http.ListenAndServe(":8080", router) // test localhost:8080/
	http.ListenAndServe(":"+port, router) // test heroku
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home endpoint accessed")
	fmt.Fprintf(w, "Aplication is running")
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllEvents endpoint accessed")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
