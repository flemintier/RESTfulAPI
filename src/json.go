package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/flemintier/RESTfulAPI/config"
)

// Fonction pour indenter la structure des données json
func IndentJson() []byte {
	file, err := json.MarshalIndent(DocConfig, "", " ")
	if err != nil {
		fmt.Println("err : ", err.Error())
		panic("Problème avec l'indentation'")
	}
	return file
}

// Fonction pour récupérer les données de fichier json
func getJson(filePath string, v any) {
	// Lecture des données du fichier json
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("err : ", err.Error())
		panic("Problème avec le fichier json")
	}

	// récupération de la structure du fichier json dans un objet
	err = json.Unmarshal(file, &v)
	if err != nil {
		fmt.Println("err : ", err.Error())
		panic("Problème avec les données de configuration")
	}
}

// Fonction de récupération du fichier json pour stocker le contenu dans la variable global "[]docs"
func getJsonConfig() config.Config {
	var data []config.Config

	// Récupération de la configuration du fichier json
	getJson("config/config.json", &data)

	return data[0]
}

// Fonction de récupération du fichier json pour stocker le contenu dans la variable global "[]docs"
func getJsonData() {
	// Récupération des documents du fichier json
	getJson("resources/documents.json", &DocConfig)
}
