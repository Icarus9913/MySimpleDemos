package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	//建立链接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if nil != err {
		panic(err)
	}
	defer conn.Close()

	//获取channel
	ch, err := conn.Channel()
	if nil != err {
		panic(err)
	}
	defer ch.Close()

	//推模式获取msgs
	msgs, err := ch.Consume("queue.dlx", "", true, false, false, false, nil)
	if nil != err {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("收到消息: %s\n", msg.Body)
		}
	}()

	fmt.Printf("[*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
