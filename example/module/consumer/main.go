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

	queueDetails.Connection = mqConn
	queueDetails.Channel = channel
	queueDetails.RouteKey = "routeKey"

	//to listen and recieve message from queue using the connection and channel
	_, deliveries := queues.Queue.Consume(queueDetails)

	forever := make(chan bool)
	go func() {
		for resp := range deliveries {
			/*
				Do something
			*/
			//Check if message comes from specific route
			if resp.RoutingKey == queueDetails.RouteKey {
				fmt.Println(string(resp.Body))

				// acknowledge the message once done with the task
				resp.Ack(false)
			}

		}
	}()

	log.Println("[-] Waiting for Messages ... ")
	<-forever
}
