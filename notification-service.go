package notification_service

const NotifierTypeMail = "mail"
const NotifierTypeRocket = "rocket"

type Notifier interface {
	SendMessage(msg string, dest string) (string, error)
	GetName() string
}

type NotifierConfig struct {
	Name   string               `mapstructure:"name"`
	Type   string               `mapstructure:"type"`
	Source NotifierSourceConfig `mapstructure:"source"`
	Host   string               `mapstructure:"host"`
	Port   int                  `mapstructure:"port"`
	Tls    bool                 `mapstructure:"tls"`
	Debug  bool                 `mapstructure:"debug"`
}

type NotifierSourceConfig struct {
	From string `mapstructure:"from"`
	Pwd  string `mapstructure:"pwd"`
}
