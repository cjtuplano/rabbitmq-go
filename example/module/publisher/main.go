package main

import (
	"fmt"
	"log"

	"github.com/cjtuplano/rabbitmq-go"
)

func main() {

	queueDetails := queues.QueueDetails{
		QueueName:    "queueName",
		ExchangeName: "exchangeName",
		ExchangeType: "direct",
	}

	//use to create connection and channel and also to declare a queue
	mqConn, channel, err := queues.Queue.ConnectMQ(queueDetails)

	if err != nil {
		log.Fatal(err)
	}

	//send the message by using the same connection and channel
	queueDetails.Message = []byte("hello world!")
	queueDetails.Connection = mqConn
	queueDetails.Channel = channel
	queueDetails.RouteKey = "routeKey"

	res := queues.Queue.Publish(queueDetails)
	fmt.Println(res)
}
