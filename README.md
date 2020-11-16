# notification-service

A library providing the ability to send messages.

This project aims to be able to send messages by **Mail** and/or by **Rocket**.
You can easily implement another notifier type, you just have to respect the interface.

### Usage

#### Straightforward usage

Create an SMTP notifier

```go
notifier, _ := mail.InitNotifier(notification_service.NotifierConfig{
	Name:   "mail-example",
	Type:   "mail",
	Source: notification_service.InfoConfSource{From: "example@gmail.com", Pwd: "password"},
	Host:   "smtp.gmail.com",
	Port:   587,
	Tls:    true,
	Debug:  false,
})
```

Create a rocket chat notifier

```go
notifier, _ := rocket.InitNotifier(notification_service.NotifierConfig{
	Name:   "rocket-example",
	Type:   "rocket",
	Source: notification_service.InfoConfSource{From: "example@gmail.com", Pwd: "password"},
	Host:   "rocket.example.io",
	Port:   443,
	Tls:    true,
	Debug:  false,
})
```

Use the notifier to send email/rocket messages

```go
// Send an email using smtp notifier
_ = notifier.SendMessage("Test message.", "example@gmail.com")

// Send a rocket msg to a user
_ = notifier.SendMessage("Test message.", "@user")

// Send a rocket msg to a channel
_ = notifier.SendMessage("Test message.", "#channel")
```

Fill the structure using [Viper](https://github.com/spf13/viper#putting-values-into-viper)

```go
func ParseConfig(configFilePath string) (*Config, error) {
	var configArray Config
	viper.SetConfigName("config")
	viper.SetConfigType(json)
	viper.AddConfigPath("path/to/config/file")
	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&configArray)
	return &configInfo, nil
}
```

#### Using a configuration file

If you want to create a configuration file, you can use [Viper](https://github.com/spf13/viper#putting-values-into-viper) to read,
and fill this structure by Unmarshalling the config file. The `mapstructure` will read all configuration file type.

```go
type Config struct {
    Notifiers []notification_service.NotifierConfig `mapstructure:"notifiers"`
}
```

Here is a JSON configuration file example

```json
"notifiers": [
    {
      "name": "smtp-example",
      "type": "mail",
      "source": {
        "from": "example@gmail.com",
        "pwd": "secret_pwd"
      },
      "host": "smtp.gmail.com",
      "port": 587,
      "tls": true,
      "debug": false
    },
    {
      "name": "rocket-example",
      "type": "rocket",
      "source": {
        "from": "example@gmail.com",
        "pwd": "secret_pwd"
      },
      "host": "rocket.example.io",
      "port": 443,
      "tls": true,
      "debug": false
    }
  ]
```
