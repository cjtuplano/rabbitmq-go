RABBITMQ Interface
Includes the following:
    - Producer/Publisher
    - Consumer
    - Queue

How to use:
For Publisher:
    - Create a connection, channel and declare a Queue using ConnectMQ
    - Publish or send a message to Publish (from package `queues`)

For Consumer:
    - Create a connection, channel and declare a queue using ConnectMQ
    - Wait for a message from Consume (from package `queues`)