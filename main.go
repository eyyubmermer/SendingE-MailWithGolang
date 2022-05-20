package main

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	abc := gomail.NewMessage()

	abc.SetHeader("From", "<sendermail>")
	abc.SetHeader("To", "<receiver>")
	abc.SetHeader("Subject", "This is an e-mail sent by Golang!")
	abc.SetBody("text/plain", "This is an e-mail sent by Golang!")

	a := gomail.NewDialer("smtp.gmail.com", 587, "<sendermail>", "<password>")

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
