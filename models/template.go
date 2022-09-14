package models

import "github.com/flemintier/RESTfulAPI/config"

type TemplateData struct {
	Docs map[string][]config.Doc
}
