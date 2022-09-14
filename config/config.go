package config

import "text/template"

type TemplateConfig struct {
	TemplateCache map[string]*template.Template
}

type Config struct {
	Port string
}

// Structure des documents
type Doc struct {
	ID          int    `json:"ID"`
	Nom         string `json:"Nom"`
	Description string `json:"Description"`
}
