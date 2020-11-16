package mail

import (
	"net/smtp"
	"strconv"

	"github.com/DumesnyJeremy/notification-service/implementations"
)

type Mail struct {
	Config implementations.NotifierConfig
}

// Receives the config extract from the configuration file and add the sender mail and password.
func InitNotifier(config implementations.NotifierConfig) (implementations.Notifier, error) {
	Notifier, err := initMail(config)
	if err != nil {
		return nil, err
	}
	return Notifier, nil
}

func initMail(notifierConfig implementations.NotifierConfig) (*Mail, error) {
	return &Mail{
		Config: notifierConfig,
	}, nil
}

// Take in argument the message and the dest and
// use smtp.SendMail and smtp.PlainAuth to send the mail
func (mail *Mail) SendMessage(msg string, dest string) (string, error) {
	to := dest
	subject := "Subject: " + msg + "\n\n"
	message := "From: " + mail.Config.Source.From + "\n" +
		"To: " + to + "\n" +
		subject +
		msg

	port := mail.Config.Host + ":" + strconv.Itoa(mail.Config.Port)
	err := smtp.SendMail(port,
		smtp.PlainAuth("", mail.Config.Source.From, mail.Config.Source.Pwd, mail.Config.Host),
		mail.Config.Source.From, []string{to}, []byte(message))
	if err != nil {
		return "", err
	}
	return "Mail", nil
}

// Return the notifier name.
func (mail *Mail) GetName() string {
	return mail.Config.Name
}
