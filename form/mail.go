package form

import (
	"fmt"
	"log"
	"net/smtp"
)

func MailToken(tokenKey string, email string) error {
	from := "biniyamdemissew112@gmail.com"
	password := `biniyam112` //'#biniyam112$'
	to := []string{email}
	message := []byte("<a href='http://localhost:8088/signup/validate?conftok=" + tokenKey + "'" + ">Activate account</a>")
	fmt.Println(string(message))
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	//sends email to the specified mail from user signin
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		to,
		message)

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
	return err
}
