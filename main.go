package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	//access log
	log.Println("Starting API...")
	//endpoints
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	// http.ListenAndServe(":8080", router) // test localhost:8080/
	http.ListenAndServe(":"+port, router) // test heroku
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home endpoint accessed")
	fmt.Fprintf(w, "Aplication is running")
}
