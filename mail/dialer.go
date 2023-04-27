package mail

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/jnnkrdb/corerdb/fnc"
	gomail "gopkg.in/mail.v2"
)

// mailobjects of this package

// mail class, contains information about the mail, receivers, etc.
type MailMessage struct {
	From       string   `json:"from"`
	To         []string `json:"to"`
	Cc         []string `json:"cc"`
	Bcc        []string `json:"bcc"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
	TextFormat string   `json:"textformat"`
	/*
		_Attachements struct {
			FileName string `json:"filename"`
			Content  string `json:"content"`
		} `json:"attachments"`
	*/
}

// smtp server class
type SMTPServer struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	// must be base64 encoded
	Password string `json:"password"`
}

// sending the mailmessage
//
// Parameters:
//   - `mailserver` : SMTPServer > information about the smtp-server connection and authentication
//   - `msg` : MailMessage > structural information about the message, which will be send
func SendMail(mailserver SMTPServer, msg MailMessage) (err error) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", msg.From)

	if len(msg.To) == 0 {
		err = fmt.Errorf("list of receivers [To] cannot be empty")
		return
	}

	mail.SetHeader("To", strings.Join(msg.To, ";"))

	switch {
	case len(msg.Cc) > 0:
		mail.SetHeader("Cc", strings.Join(msg.Cc, ";"))
	case len(msg.Bcc) > 0:
		mail.SetHeader("Bcc", strings.Join(msg.Bcc, ";"))
	}

	mail.SetHeader("Subject", msg.Subject)
	if fnc.StringInList(msg.TextFormat, []string{
		"text/html",
		"text/plain",
		"text/richtext",
		"text/rtf",
		"text/xml",
	}) {
		mail.SetBody(msg.TextFormat, msg.Body)
	} else {
		mail.SetBody("text/plain", msg.Body)
	}

	// --- placeholder attachments

	// mail.Attach()

	dialer := gomail.NewDialer(
		mailserver.Address,
		mailserver.Port,
		mailserver.Username,
		fnc.UnencodeB64(mailserver.Password),
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = dialer.DialAndSend(mail)

	return
}
