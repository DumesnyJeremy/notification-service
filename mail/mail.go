package mail

import (
	"net/smtp"
	"strconv"

	"ssl-ar/notification-service"
)

type Mail struct {
	Config notification_service.NotifierConfig
	from   string
	pass   string
}

// Receives the config extract from the configuration file and add the sender mail and password.
func InitNotifier(config notification_service.NotifierConfig) (notification_service.Notifier, error) {
	Notifier, err := initMail(config)
	if err != nil {
		return nil, err
	}
	return Notifier, nil
}

func initMail(notifierConfig notification_service.NotifierConfig) (*Mail, error) {
	from := notifierConfig.Source.From
	pass := notifierConfig.Source.Pwd
	return &Mail{
		Config: notifierConfig,
		from:   from,
		pass:   pass,
	}, nil
}

// Take in argument the message and the dest and
// use smtp.SendMail and smtp.PlainAuth to send the mail
func (mail *Mail) SendMessage(msg string, dest string) (string, error) {
	to := dest
	subject := "Subject: " + msg + "\n\n"
	message := "From: " + mail.from + "\n" +
		"To: " + to + "\n" +
		subject +
		msg

	port := mail.Config.Host + ":" + strconv.Itoa(mail.Config.Port)
	err := smtp.SendMail(port,
		smtp.PlainAuth("", mail.from, mail.pass, mail.Config.Host),
		mail.from, []string{to}, []byte(message))
	if err != nil {
		return "", err
	}
	return "Mail", nil
}

// Return the notifier name.
func (mail *Mail) GetName() string {
	return mail.Config.Name
}
