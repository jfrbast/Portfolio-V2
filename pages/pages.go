package pages

import (
	"Portfolio/mailer"
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
func ProjectsPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Temp.ExecuteTemplate(w, "projects", nil)
	if err != nil {
		log.Println("Erreur template accueil:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	Name := r.FormValue("name")
	Email := r.FormValue("email")
	Message := r.FormValue("message")

	var msg, msgType string

	success := mailer.SendMail(Email, Name, Message)
	if success == "Missing" {
		msg = "Please fill in all fields"
		msgType = "error"
	}
	if success == "Sent" {
		msg = "Message sent successfully"
		msgType = "success"
	}
	if success == "Error" {
		msg = "Error sending message"
		msgType = "error"
	}

	err := templates.Temp.ExecuteTemplate(w, "contact", map[string]interface{}{
		"Message":     msg,
		"MessageType": msgType,
	})
	if err != nil {
		log.Println("Erreur template contact:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}
