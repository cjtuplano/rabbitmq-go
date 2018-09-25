package main

import (
	"fmt"
	"log"
	queues "rabbitmq-go/module/queues/services"
)

func main() {
	//use to create connection and channel and also to declare a queue
	mqConn, channel, err := queues.Queue.ConnectMQ(queues.QueueListener{
		QueueName:    "sample1",
		ExchangeName: "exchangeName",
		ExchangeType: "direct",
	})

	if err != nil {
		log.Fatal(err)
	}

	//send the message by using the same connection and channel
	publisherData := queues.QueueListener{QueueName: "sample1", Message: []byte("hello"), Connection: mqConn, Channel: channel}

	res := queues.Queue.Publish(publisherData)
	fmt.Println(res)
}
