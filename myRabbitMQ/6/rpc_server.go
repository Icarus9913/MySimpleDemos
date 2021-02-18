package main

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

//go run rpc_server.go

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

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

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if nil != err {
		panic("Failed to declare a queue: " + err.Error())
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if nil != err {
		panic("Failed to set QoS: " + err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if nil != err {
		panic("Failed to register a consumer: " + err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			if nil != err {
				panic("Failed to convert body to integer: " + err.Error())
			}

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			if nil != err {
				panic("Failed to publish a message: " + err.Error())
			}

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
