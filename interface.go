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

//QueueDetails - data structure
type QueueDetails struct {
	ExchangeName string           `json:"exchangeName"`
	ExchangeType string           `json:"exchangeType"`
	BindKey      string           `json:"bindingKey"`
	QueueName    string           `json:"queueName"`
	Message      []byte           `json:"message"`
	Connection   *amqp.Connection `json:"connection"`
	Channel      *amqp.Channel    `json:"channel"`
	RouteKey     string           `json:"routeKey"`
}

//ConnectMQ - for one time rabbitmq connection
func (queueDetails QueueDetails) ConnectMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(settings.MQSettings.Link)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open Channel")

	// declare an exchange
	/*
		ExchangeDeclare arguments (name, kind string, durable, autoDelete, internal, noWait bool, args Table)
	*/

	err = ch.ExchangeDeclare(
		queueDetails.ExchangeName,
		queueDetails.ExchangeType,
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
		queueDetails.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	err = ch.QueueBind(
		queueDetails.QueueName,    // queue name
		queueDetails.RouteKey,     // routing key
		queueDetails.ExchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	return conn, ch, err
}

//Publish - Queue Publisher
func (queueDetails QueueDetails) Publish() string {

	ch := queueDetails.Channel

	err := ch.Publish(
		queueDetails.ExchangeName, //exchange name
		queueDetails.RouteKey,     // route key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         queueDetails.Message,
		})
	log.Printf(" [x] Sent %s", string(queueDetails.Message))
	failOnError(err, "Failed to publish a message")
	return "Message sent"
}

//Consume - Queue Consumer
func (queueDetails QueueDetails) Consume() (*amqp.Connection, <-chan amqp.Delivery) {

	conn := queueDetails.Connection
	ch := queueDetails.Channel

	err := ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	deliveries, err := ch.Consume(
		queueDetails.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	return conn, deliveries

}
