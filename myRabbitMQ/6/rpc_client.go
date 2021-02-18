package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

/*
	将使用RabbitMQ构建一个RPC系统：客户端和可伸缩RPC服务器。
	由于我们没有值得分配的耗时任务，因此我们将创建一个虚拟RPC服务，该服务返回斐波那契数。

	go run rpc_client.go 30
*/

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
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
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if nil != err {
		panic("Failed to declare a queue: " + err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if nil != err {
		panic("Failed to register a consumer: " + err.Error())
	}

	corrId := randomString(32)

	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	if nil != err {
		panic("Failed to publish a message: " + err.Error())
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			if nil != err {
				panic("Failed to convert body to integer: " + err.Error())
			}
			break
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := bodyFrom(os.Args)

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(n)
	if nil != err {
		panic("Failed to handle RPC request: " + err.Error())
	}

	log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	if nil != err {
		panic("Failed to convert arg to integer: " + err.Error())
	}
	return n
}
