package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/flemintier/RESTfulAPI/config"
	"github.com/flemintier/RESTfulAPI/models"
)

// Déclarations de variables
var appTmpl *config.TemplateConfig

// Fonction d'initialisation de la variable grobales du template
func CreateTemplates(app *config.TemplateConfig) {
	appTmpl = app
}

// Fonction qui va récupérer et exécuter les templates et afficher les données sur la page web
func RenderTemplate(rw http.ResponseWriter, tmplName string, td *models.TemplateData) {
	templateCache := appTmpl.TemplateCache

	// templateCache["home.page.tmpl"]
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(rw, "Le template "+tmplName+" n'existe pas!", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	// fmt.Println(*td)
	err := tmpl.Execute(buffer, td)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution du template : ", err.Error())
		return
	}
	// fmt.Println(buffer.String())
	buffer.WriteTo(rw)
	json.NewEncoder(rw).Encode(*td)
}

// Fonction de gestion du cache pour que le site s'affiche correctement
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layout, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layout) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
