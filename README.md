#notification-service

A library providing the ability to send messages.

This project aims to be able to send messages by **Mail** and/or by **Rocket**.
You can easily implement another notifier type, you just have to respect the interface.

### Usage
#### With Configuration file
```json
"notifiers": [
    {
      "name": "gmail-example",
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
If you want to create a configuration file, you can use [Viper](https://github.com/spf13/viper#putting-values-into-viper) to read,
and fill this structure by Unmarshalling the config file. The `mapstructure` will read all configuration file type.
```go
type Config struct {
    Notifiers []notification_service.NotifierConfig `mapstructure:"notifiers"`
}
```

#### Without configuration file
Mail example
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
	_, _ = notifier.SendMessage("Test message.", "example@gmail.com")
```
Rocket example
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
	_, _ = notifier.SendMessage("Test message.", "@example")
```
