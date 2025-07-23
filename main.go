package main

import (
	"Portfolio/pages"
	"Portfolio/templates"
	"fmt"
	"log"
	"net/http"
)

func main() {
	templates.InitTemplates()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", pages.HomePage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
