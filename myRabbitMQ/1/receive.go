package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if nil!=err{
		panic("Failed to connect to RabbitMQ: "+err.Error())
	}
	defer conn.Close()

	// 获取channel
	ch, err := conn.Channel()
	if nil!=err{
		panic("Failed to open a channel: "+err.Error())
	}
	defer ch.Close()

	//定义队列的长度
	args := make(amqp.Table)
	args["x-max-length"] = 1000

	// 声明队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		args,     // arguments
	)
	if nil!=err{
		panic("Failed to declare a queue: "+err.Error())
	}
	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if nil!=err{
		panic("Failed to register a consumer: "+err.Error())
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
