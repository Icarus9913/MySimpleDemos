package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

/*
	如果要将日志保存到文件，只需打开控制台并输入：go run receive_logs.go > logs_from_rabbit.log
	如果希望在屏幕上查看日志，请切换到一个新的终端并运行：go run receive_logs.go

*/

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("connect to RabbitMQ failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("open a channel failed, err:%v\n", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
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

	//我们已经创建了一个扇出交换器和一个队列。现在我们需要告诉交换器将消息发送到我们的队列。
	//交换器和队列之间的关系称为*绑定*。
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil)
	if nil != err {
		panic("Failed to bind a queue: " + err.Error())
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
