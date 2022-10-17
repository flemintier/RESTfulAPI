package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flemintier/RESTfulAPI/config"
	"github.com/gorilla/mux"
)

// Déclarations des variables
var appConfig config.Config
var DocConfig map[string]config.Doc

// Requète des pages
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getAllDocs).Methods("GET")
	myRouter.HandleFunc("/{Nom}", getDocs).Methods("GET")
	myRouter.HandleFunc("/post", postDocs).Methods("POST")
	myRouter.HandleFunc("/{Nom}", deleteDocs).Methods("DELETE")
	log.Fatal(http.ListenAndServe(appConfig.Port, myRouter))
}

// Début du programme
func main() {
	// Récupération des données de configuration
	appConfig = getJsonConfig()
	getJsonData()

	// Page Home
	fmt.Println("(http://localhost:8080) - Server started on port", appConfig.Port)

	handleRequests()
}
