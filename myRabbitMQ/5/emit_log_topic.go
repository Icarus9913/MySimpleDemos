package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

/*
	topic交换器

	并发出带有路由键“kern.critical”的日志：	go run emit_log_topic.go "kern.critical" "A critical kernel error"
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

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic",          // exchange
		severityFrom(os.Args), // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if nil != err {
		panic("Failed to publish a message: " + err.Error())
	}

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
