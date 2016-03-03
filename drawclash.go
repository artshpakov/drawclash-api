package main

import (
	"github.com/artshpakov/drawclash/app/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", api.UsersShowAction)

	http.Handle("/", router)
	log.Print("Server running at port 8080\n")
	http.ListenAndServe(":8080", nil)
}
