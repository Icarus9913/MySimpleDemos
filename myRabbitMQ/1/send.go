package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1. 尝试连接RabbitMQ，建立连接, 该连接抽象了套接字socket连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//定义队列的长度
	args := make(amqp.Table)
	args["x-max-length"] = 1000

	// 3. 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		args,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"

	// 4.将消息发布到声明的队列
	err = ch.Publish(
		"",     // exchange	交换器,此时为默认/无名
		q.Name, // routing key			消息将以`route_key`参数指定的名称路由到队列
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
