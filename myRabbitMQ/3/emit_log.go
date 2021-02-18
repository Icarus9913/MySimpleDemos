package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

/*
	列出服务器上的交换器: sudo rabbitmqctl list_exchanges

	将向多个消费者传递一个消息。这就是所谓的“订阅/发布模式”。
		*生产者*是发送消息的用户应用程序。 	provider
		*队列*是存储消息的缓冲区。 			queue
		*消费者*是接收消息的用户应用程序。	consumer
	生产者只能将消息发送到"交换器"(exchanges)。交换器是非常简单的东西。一方面，它接收来自生产者的消息，
									另一方面，将它们推入队列。交换器必须确切知道如何处理接收到的消息。
	{有几种交换器类型可用：direct, topic, headers 和 fanout.}
*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name		使用命名的交换器
		"fanout", // type		交换器类型
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}


/*
临时队列:
	在amqp客户端中，当我们传递一个空字符串作为队列名称时，我们将使用随机生成的名称创建一个非持久队列,
	一旦我们断开消费者的连接，队列就会自动删除:

	q, err := ch.QueueDeclare(
	  "",    // 空字符串作为队列名称
	  false, // 非持久队列
	  false, // delete when unused
	  true,  // 独占队列（当前声明队列的连接关闭后即被删除）
	  false, // no-wait
	  nil,   // arguments
	)

上述方法返回时，生成的队列实例包含RabbitMQ生成的随机队列名称。例如，它可能看起来像amq.gen-JzTY20BRgKO-HjmUJj0wLg。

*/