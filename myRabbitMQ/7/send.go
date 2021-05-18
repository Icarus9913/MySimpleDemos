package main

import "github.com/streadway/amqp"

func main() {
	//1.创建rabbitmq客户端socket链接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if nil != err {
		panic(err)
	}
	defer conn.Close()

	//2.创建channel
	ch, err := conn.Channel()
	if nil != err {
		panic(err)
	}
	defer ch.Close()

	//声明死信交换器
	err = ch.ExchangeDeclare("exchange.dlx", "direct", true, false, false, false, nil)
	if nil != err {
		panic(err)
	}
	//声明正常的交换器
	err = ch.ExchangeDeclare("exchange.normal", "fanout", true, false, false, false, nil)
	if nil != err {
		panic(err)
	}

	//配置参数
	args := make(amqp.Table)
	args["x-message-ttl"] = 5000 //5000毫秒,也就是5秒钟
	args["x-dead-letter-exchange"] = "exchange.dlx"
	args["x-dead-letter-routing-key"] = "routingkey"

	//声明普通的queue
	normalQueue, err := ch.QueueDeclare("queue.normal", true, false, false, false, args)
	if nil != err {
		panic(err)
	}
	//将普通的queue与normal交换器绑定
	err = ch.QueueBind(normalQueue.Name, "", "exchange.normal", false, nil)
	if nil != err {
		panic(err)
	}

	//声明死信队列
	dlxQueue, err := ch.QueueDeclare("queue.dlx", true, false, false, false, nil)
	if nil != err {
		panic(err)
	}
	//将死信队列与死信交换器绑定,其中BindingKey是"routingkey"
	err = ch.QueueBind(dlxQueue.Name, "routingkey", "exchange.dlx", false, nil)
	if nil != err {
		panic(err)
	}

	body := "hahahhhh"
	//将消息发布到声明的正常队列中
	for i := 0; i < 10; i++ {
		err = ch.Publish(
			"exchange.normal",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if nil != err {
			panic(err)
		}
	}
}
