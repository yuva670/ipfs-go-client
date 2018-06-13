package email

import (
	"crypto/tls"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

func SendMail(toEmail string, fromEmail string, body string) {
	// Set up authentication information.
	servername := "smtp.gmail.com:465"
	host, _, _ := net.SplitHostPort(servername)

	from := mail.Address{"", fromEmail}
	to := mail.Address{"", toEmail}
	auth := smtp.PlainAuth("", "abc@def", "123456", host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte(
		"Subject: File shared with you\r\n" +
			"\r\n" +
			body + "\r\n")
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
