package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flemintier/RESTfulAPI/config"
	"github.com/flemintier/RESTfulAPI/internal/handlers"
	"github.com/flemintier/RESTfulAPI/models"
	"github.com/gorilla/mux"
)

// Déclarations des variables
var appConfig config.Config
var DocConfig *[]config.Doc

// Page d'accueil
func Home(rw http.ResponseWriter, r *http.Request) {
	// docs := make(map[string][]config.Doc)
	// docs["docs"] = DocConfig
	// fmt.Println("", *DocConfig)
	doc := make(map[string][]config.Doc)
	doc["docs"] = *DocConfig
	handlers.RenderTemplate(rw, "home", &models.TemplateData{
		Docs: doc,
	})

	// d, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Oops", http.StatusBadRequest)
	// 	return
	// }
	// fmt.Fprintf(rw, "Hello %s", d)
}

// Requète des pages
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Home)
	myRouter.HandleFunc("/{Nom}", getDocs).Methods("GET")
	myRouter.HandleFunc("/post", postDocs).Methods("POST")
	myRouter.HandleFunc("/{Nom}", deleteDocs).Methods("DELETE")
	log.Fatal(http.ListenAndServe(appConfig.Port, myRouter))
}

// Début du programme
func main() {
	// Variables de configuration
	var appTmpl config.TemplateConfig

	// Récupération des données de configuration
	appConfig = getJsonConfig()
	getJsonData()

	// Indentation des données json
	// file := IndentJson()
	// fmt.Println("docs : ", docs)
	// fmt.Println("file : ", string(file))

	// ------ Site web ------ //
	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		panic(err)
	}

	appTmpl.TemplateCache = templateCache
	handlers.CreateTemplates(&appTmpl)

	// Page Home
	// http.HandleFunc("/", handlers.Home)
	fmt.Println("(http://localhost:8080) - Server started on port", appConfig.Port)
	// http.ListenAndServe(appConfig.Port, nil)

	handleRequests()
}
