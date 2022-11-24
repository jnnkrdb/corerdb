package mail

import (
	"crypto/tls"
	"strings"

	"github.com/jnnkrdb/corerdb/fnc"
	"github.com/jnnkrdb/corerdb/prtcl"
	gomail "gopkg.in/mail.v2"
)

// sending the mailmessage
//
// Parameters:
//   - `mailserver` : SMTPServer > information about the smtp-server connection and authentication
//   - `msg` : MailMessage > structural information about the message, which will be send
func SendMail(mailserver SMTPServer, msg MailMessage) (err error) {

	prtcl.Log.Println("starting mailflow")

	mail := gomail.NewMessage()

	mail.SetHeader("From", msg.From)

	mail.SetHeader("To", strings.Join(msg.To, ";"))

	if len(msg.Cc) > 0 {
		mail.SetHeader("Cc", strings.Join(msg.Cc, ";"))
	}

	if len(msg.Bcc) > 0 {
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

	if err = dialer.DialAndSend(mail); err != nil {

		prtcl.Log.Println("error sending mailmessage:", err)
	}

	return
}
