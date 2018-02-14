package main

import "log"
import "net/smtp"

func main() {

	to := []string{"ec2-user"}
	msg := []byte(
		"To: ec2-user\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err := smtp.SendMail("localhost:25", nil, "ec2-user", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

