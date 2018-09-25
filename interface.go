package queues

import (
	"log"

	"github.com/cjtuplano/rabbitmq-go/config"

	"github.com/streadway/amqp"
)

var settings = config.GetConfig()

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Queue interface that can publish and consume queues
type Queue interface {
	Publish() string
	Consume() (*amqp.Connection, <-chan amqp.Delivery)
	ConnectMQ() (*amqp.Connection, *amqp.Channel, error)
}

type QueueListener struct {
	ExchangeName string           `json:"exchangeName"`
	ExchangeType string           `json:"exchangeType"`
	BindKey      string           `json:"bindingKey"`
	QueueName    string           `json:"queueName"`
	Message      []byte           `json:"message"`
	Connection   *amqp.Connection `json:"connection"`
	Channel      *amqp.Channel    `json:"channel"`
}

//ConnectMQ - for one time rabbitmq connection
func (listener QueueListener) ConnectMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(settings.MQSettings.Link)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open Channel")

	// declare an exchange
	/*
		ExchangeDeclare arguments (name, kind string, durable, autoDelete, internal, noWait bool, args Table)
	*/

	err = ch.ExchangeDeclare(
		listener.ExchangeName,
		listener.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	//declare queue
	//QueueDeclare arguments (name string, durable, autoDelete, exclusive, noWait bool, args Table)
	_, err = ch.QueueDeclare(
		listener.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	return conn, ch, err
}

//Publish - Queue Publisher
func (listener QueueListener) Publish() string {

	ch := listener.Channel

	err := ch.Publish(
		listener.ExchangeName, //exchange type
		listener.QueueName,    // route key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         listener.Message,
		})
	log.Printf(" [x] Sent %s", string(listener.Message))
	failOnError(err, "Failed to publish a message")
	return "Message sent"
}

//Consume - Queue Consumer
func (listener QueueListener) Consume() (*amqp.Connection, <-chan amqp.Delivery) {

	conn := listener.Connection
	ch := listener.Channel

	err := ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	deliveries, err := ch.Consume(
		listener.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	return conn, deliveries

}
