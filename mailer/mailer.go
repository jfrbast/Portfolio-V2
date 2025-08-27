package mailer

import (
	"net/smtp"
)

func SendMail(From string, Name string, Message string) string {

	To := []string{"bastienjouffre@gmail.com"}
	from := "bastienjouffre@gmail.com"
	password := "rendu infonctionnel pour pas divulger mon mot de passe" //os.Getenv("GMAIL_APP_PASSWORD")

	subject := "Mail reçu du Portfolio \n"
	body := Message
	message := []byte(subject + "\n" + Name + "\n" + From + "\n" + body)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if Name == "" || Message == "" || From == "" {
		return "Missing"
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, To, message)
	if err != nil {
		return "Error"
	}
	return "Sent"

}
