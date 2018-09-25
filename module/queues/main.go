package main

import (
	"fmt"
	"log"
	"rabbitmq-go/module/queues/services"
)

type Publisher struct {
	Name string
}

//for service queues
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

	//Consume message
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
