package pages

import (
	"Portfolio/templates"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := templates.Temp.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Println("Erreur template accueil:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}
