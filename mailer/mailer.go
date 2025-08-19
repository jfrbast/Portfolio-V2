package mailer

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)


func SendMail(From string, Name string, Message string) {

	To := []string{"bastienjouffre@gmail.com"}
	from := "bastienjouffre@gmail.com"
	password := os.Getenv("GMAIL_APP_PASSWORD")

	subject := "Mail reçu du Portfolio \n"
	body := Message
	message := []byte(subject + "\n" + Name + "\n" + From + "\n" + body)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	fmt.Println("ok")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	fmt.Println("ok")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, To, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
	log.Println("✅ Email envoyé avec succès !")
}
