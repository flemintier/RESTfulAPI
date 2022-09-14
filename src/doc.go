package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flemintier/RESTfulAPI/config"
	"github.com/flemintier/RESTfulAPI/internal/handlers"
	"github.com/flemintier/RESTfulAPI/models"
	"github.com/gorilla/mux"
)

// Fonction de création de document et ajout à la liste globale des documents
func Creation(nom string, description string) config.Doc {
	iD := (*DocConfig)[len(*DocConfig)-1].ID + 1 // incrémentation de l'ID

	// Création du nouveau document
	d := config.Doc{
		ID:          iD,
		Nom:         nom,
		Description: description,
	}

	return d
}

// Fonction de récupération de document
func Recuperation(nom string) config.Doc {
	var ret config.Doc

	for _, v := range *DocConfig {
		if v.Nom == nom {
			return v
		}
	}

	return ret
}

// Fonction de suppression de document
func Suppression(nom string) {
	for i, v := range *DocConfig {
		if v.Nom == nom {
			*DocConfig = append((*DocConfig)[:i], (*DocConfig)[i+1:]...)
			break
		}
	}
}

// Fonction pour créer un document
func postDocs(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var doc config.Doc

	_ = json.NewDecoder(r.Body).Decode(&doc)
	fmt.Println("new doc : ", doc)
	// Création du nouveau document
	doc = Creation(doc.Nom, doc.Description)

	// json.NewEncoder(rw).Encode(doc)
	// Ajout du nouveau document à la liste globale
	*DocConfig = append(*DocConfig, doc)
	fmt.Println(*DocConfig)
	json.NewEncoder(rw).Encode(*DocConfig)
	// docs := make(map[string][]config.Doc)
	// docs["docs"] = DocConfig
	// handlers.RenderTemplate(rw, "home", DocConfig)
}

// Fonction pour récupérer un document
func getDocs(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)

	doc := Recuperation(params["Nom"])
	json.NewEncoder(rw).Encode(doc)
	// docs := make(map[string][]config.Doc)
	// docs["docs"] = DocConfig
	// handlers.RenderTemplate(rw, "home", DocConfig)
}

// Fonction pour supprimer un document
func deleteDocs(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)

	Suppression(params["Nom"])
	d := make(map[string][]config.Doc)
	d["docs"] = *DocConfig
	handlers.RenderTemplate(rw, "home", &models.TemplateData{
		Docs: d,
	})
}
