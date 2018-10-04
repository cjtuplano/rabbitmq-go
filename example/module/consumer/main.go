package main

import (
	"fmt"
	"log"

	"github.com/cjtuplano/rabbitmq-go"
)

func main() {
	var routeKeys = []string{
		"routeKey",
	}

	queueDetails := queues.QueueDetails{
		QueueName:    "queueName",
		ExchangeName: "exchangeName",
		ExchangeType: "direct",
		RouteKeys:    routeKeys,
	}

	//use to create connection and channel and also to declare a queue
	mqConn, channel, err := queues.Queue.ConnectMQ(queueDetails)

	if err != nil {
		log.Fatal(err)
	}

	queueDetails.Connection = mqConn
	queueDetails.Channel = channel

	//to listen and recieve message from queue using the connection and channel
	_, deliveries := queues.Queue.Consume(queueDetails)

	forever := make(chan bool)
	go func() {
		for resp := range deliveries {
			//Check if message comes from specific route
			if resp.RoutingKey == "routeKey" {
				/*
					Do Something
				*/
				fmt.Println(string(resp.Body))
				resp.Ack(false)
			}

		}
	}()

	log.Println("[-] Waiting for Messages ... ")
	<-forever
}
