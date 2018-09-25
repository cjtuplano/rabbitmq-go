package main

import (
	"fmt"
	"log"

	"github.com/cjtuplano/rabbitmq-go"
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
	//to listen and recieve message from queue using the connection and channel
	_, deliveries := queues.Queue.Consume(queues.QueueListener{
		QueueName:  "sample1",
		Connection: mqConn,
		Channel:    channel,
	})

	forever := make(chan bool)
	go func() {
		for resp := range deliveries {
			/*
				Do something
			*/

			fmt.Println(string(resp.Body))

			// acknowledge the message once done with the task
			resp.Ack(false)
		}
	}()
	<-forever
}
