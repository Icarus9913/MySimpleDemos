package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

/*
	将使它能够只订阅消息的一个子集. 例如，我们将只能将关键错误消息定向到日志文件（以节省磁盘空间），同时仍然能够在控制台上打印所有日志消息。

	go run emit_log_direct.go error "Run. Run. Or it will explode."    ->这里的双引号一整大串,只算一个字符串

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
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if nil != err {
		panic("Failed to declare an exchange: " + err.Error())
	}

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_direct",         // exchange
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
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}
