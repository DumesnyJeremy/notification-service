package rocket

import (
	"github.com/titandc/gorocket/api"
	"github.com/titandc/gorocket/rest"
	"strconv"

	"ssl-ar/notification-service"
)

type RocketClient interface {
	Send(channel *api.Channel, msg string) (*rest.PostMessageReturn, error)
	GetMessages(channel *api.Channel, page *rest.Page) ([]api.Message, error)
}

type Rocket struct {
	Config notification_service.NotifierConfig
	Client RocketClient
}

// Receives the config extract from the configuration file, create the rest.NewClient and login a user
// with the Email and the Password and return a Notifier object fill in initRocket.
func InitNotifier(config notification_service.NotifierConfig) (notification_service.Notifier, error) {
	Notifier, err := initRocket(config)
	if err != nil {
		return nil, err
	}
	return Notifier, nil
}

func initRocket(notifierConfig notification_service.NotifierConfig) (*Rocket, error) {
	Client, err := initClient(notifierConfig)
	if err != nil {
		return nil, err
	}
	return &Rocket{
		Config: notifierConfig,
		Client: Client,
	}, nil
}

func initClient(config notification_service.NotifierConfig) (*rest.Client, error) {
	client := rest.NewClient(config.Host, strconv.Itoa(config.Port), config.Tls, config.Debug)
	if err := client.Login(api.UserCredentials{Email: config.Source.From, Password: config.Source.Pwd}); err != nil {
		return nil, err
	}

	if _, err := client.GetPublicChannels(); err != nil {
		return nil, err
	}
	return client, nil
}

// Take in argument the message and the dest and
// send a message to a channel or via Client.Send .
func (rocket *Rocket) SendMessage(msg string, dest string) (string, error) {
	if _, err := rocket.Client.Send(&api.Channel{Name: dest}, msg); err != nil {
		return "", err
	}
	_, _ = rocket.Client.GetMessages(&api.Channel{Name: dest}, nil)
	return "Rocket", nil
}

// Return the notifier name.
func (rocket *Rocket) GetName() string {
	return rocket.Config.Name
}
