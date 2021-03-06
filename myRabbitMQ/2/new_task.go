package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

/*
	工作队列,用于在多个工人之间分配耗时的任务:

	其中可运行多个worker.如果我们的工作正在积压，我们可以增加更多的工人，这样就可以轻松扩展。
*/


func main() {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("connect to RabbitMQ failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("open a channel failed, err:%v\n", err)
		return
	}
	defer ch.Close()

	// 3. 要发送，我们必须声明要发送到的队列。
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // 持久的
		false,        // delete when unused
		false,        // 独有的
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		fmt.Printf("declare a queue failed, err:%v\n", err)
		return
	}

	// 4. 然后我们可以将消息发布到声明的队列
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     // exchange	交换器,此时为默认/无名
		q.Name, // routing key			消息将以`route_key`参数指定的名称路由到队列
		false,  // 立即
		false,  // 强制
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 持久
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		fmt.Printf("publish a message failed, err:%v\n", err)
		return
	}
	log.Printf(" [x] Sent %s", body)
}

// bodyFrom 从命令行获取将要发送的消息内容
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}