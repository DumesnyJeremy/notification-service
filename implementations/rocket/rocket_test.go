package rocket

import (
	"github.com/titandc/gorocket/api"
	"github.com/titandc/gorocket/rest"
	"testing"

	"github.com/DumesnyJeremy/notification-service/implementations"
)

type Client struct{}

func (m *Client) Send(channel *api.Channel, msg string) (*rest.PostMessageReturn, error) {
	// Change the return to check how react Rocket implementation.
	return &rest.PostMessageReturn{}, nil
}

func (m *Client) GetMessages(channel *api.Channel, page *rest.Page) ([]api.Message, error) {
	// Change the return to check how react Rocket implementation.
	return []api.Message{}, nil
}

func TestClient(t *testing.T) {
	mockedClientObj := new(Client)
	infoRocket := fillRocketStruct(mockedClientObj)
	if _, err := infoRocket.SendMessage("Hello word", "@example"); err != nil {
		t.Error("Error: ", err)
	}
}

func fillRocketStruct(mockedClientObj *Client) Rocket {
	return Rocket{
		Config: implementations.NotifierConfig{
			Name: "rocket-example",
			Type: "rocket",
			Source: implementations.InfoConfSource{
				From: "example@gmail.com",
				Pwd:  "secret_pwd",
			},
			Host:  "rocket.example.io",
			Port:  443,
			Tls:   true,
			Debug: false,
		},
		Client: mockedClientObj,
	}
}
