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

func AboutPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Temp.ExecuteTemplate(w, "about", nil)
	if err != nil {
		log.Println("Erreur template accueil:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}
func CreditsPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Temp.ExecuteTemplate(w, "credits", nil)
	if err != nil {
		log.Println("Erreur template accueil:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}
