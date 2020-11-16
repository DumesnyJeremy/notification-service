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
You can directly init the 2 implementation, by filing the structure argument of the inits
Mail example
```go
notifier := mail.InitNotifier(notification_service.NotifierConfig{
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
    }
})
notifier.SendMessage("Test message.", example@gmail.com)
```
Rocket example
```go
notifier := rocket.InitNotifier(notification_service.NotifierConfig{
      "name": "rocket-example",
      "type": "rocket",
      "source": {
        "from": "example.example@gmail.com",
        "pwd": "secret_pwd"
      },
      "host": "rocket.example.io",
      "port": 443,
      "tls": true,
      "debug": false
})
notifier.SendMessage("Test message.", @example)
```
