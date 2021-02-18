package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

/*
	想要接收所有的日志：	go run receive_logs_topic.go "#"
	要从“kern”接收所有日志：	go run receive_logs_topic.go "kern.*"
	如果你只想接收“critical”日志：	go run receive_logs_topic.go "*.critical"
	可以创建多个绑定：	go run receive_logs_topic.go "kern.*" "*.critical"
*/

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if nil != err {
		panic("Failed to connect to RabbitMQ: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if nil != err {
		panic("Failed to open a channel: " + err.Error())
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if nil != err {
		panic("Failed to declare an exchange: " + err.Error())
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if nil != err {
		panic("Failed to declare a queue: " + err.Error())
	}

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}

	// 绑定topic
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil)
		if nil != err {
			panic("Failed to bind a queue: " + err.Error())
		}
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if nil != err {
		panic("Failed to register a consumer: " + err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}


