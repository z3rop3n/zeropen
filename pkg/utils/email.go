package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, content string, from string, password string) error {
	fmt.Println("Sending email to", to)

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Auth OTP\n\n" +
		content

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}
