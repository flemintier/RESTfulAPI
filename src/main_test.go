package main

import (
	"testing"
)

// Test de création de document
func TestCreation(t *testing.T) {
	nom := "nom"
	description := "description"
	newDoc := Creation(nom, description)

	if newDoc.Nom != nom {
		t.Error("Le nom du nouveau document n'est pas le bon, valeur : " + newDoc.Nom + ", attendue : " + nom)
	}
	if newDoc.Description != description {
		t.Error("La description du nouveau document n'est pas la bonne, valeur : " + newDoc.Description + ", attendue : " + description)
	}
}

// Test de récupération de document
func TestRecuperation(t *testing.T) {
	nom := "nom"
	doc := Recuperation(nom)

	if doc.Nom != nom {
		t.Error("Le nom du document récupéré n'est pas le bon, valeur : " + doc.Nom + ", attendue : " + nom)
	}
}

// Test de suppression de document
func TestSuppression(t *testing.T) {
	nom := "nom"
	Suppression(nom)

	for _, v := range DocConfig {
		if v.Nom == nom {
			t.Error("Le nom du document supprimé es toujours présent")
		}
	}
}
