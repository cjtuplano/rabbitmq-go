package queuemodel

import (
	"github.com/streadway/amqp"
)

type Publisher struct {
	QueueName  string           `json:"queueName"`
	Message    []byte           `json:"message"`
	Connection *amqp.Connection `json:"connection"`
	Channel    *amqp.Channel    `json:"channel"`
}
