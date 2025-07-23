package templates

import (
	"html/template"
	"log"
)

var Temp *template.Template

func InitTemplates() {
	var err error
	Temp, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du chargement des templates :", err)
	}
}
