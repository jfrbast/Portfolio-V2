package pages

import (
	"Portfolio/mailer"
	"Portfolio/templates"
	"fmt"
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

func ContactPage(w http.ResponseWriter, r *http.Request) {

	Name := r.FormValue("name")
	Email := r.FormValue("email")
	Message := r.FormValue("message")
	fmt.Println(Name, Email, Message)

	if Name != "" && Email != "" && Message != "" {
		mailer.SendMail(Email, Name, Message)
		log.Println("✅ Email envoyé avec succès !")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := templates.Temp.ExecuteTemplate(w, "contact", nil)
	if err != nil {
		log.Println("Erreur template accueil:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)

	}
}
