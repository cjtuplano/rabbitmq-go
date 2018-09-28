package config

import (
	"log"

	"rabbitmq-go/config/model"
)

//GetConfig function
func GetConfig() configmodel.ConfigSettings {
	env := GetEnv()

	settings := configmodel.ConfigSettings{}

	switch env {
	case "dev":
		settings = configmodel.ConfigSettings{
			MQSettings: configmodel.MQSettings{
				Link: "amqp://guest:guest@127.0.0.1:5672/",
			}}

	case "prod":
		settings = configmodel.ConfigSettings{
			MQSettings: configmodel.MQSettings{
				Link: "amqp://guest:guest@127.0.0.1:5672/",
			}}

	default:
		log.Fatal("Configuration not found")
	}

	return settings
}
