package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

/*
	生产者生产消息确认有两种方式: 1.事务	2.发送方确认机制

	生产者将信道设置成confirm模式,一旦信道进入confirm模式,所有在该信道上面发布的消息都会被指派一个唯一的ID(从1开始),
	一旦消息被投递到所有匹配的队列之后,RabbitMQ就会发送一个确认(Basic.Ack)给生产者(包含消息的唯一ID).
	如果消息和队列是可持久化的,那么确认消息会在消息写入磁盘之后发出.RabbitMQ回传给生产者的确认消息中的deliveryTag包含了确认消息的序号,
	此外RabbitMQ也可以设置channel.basicAck方法中的multiple参数,表示到这个序号之前的所有消息都已经得到了处理.
*/

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if nil != err {
		panic(err)
	}

	ch, err := conn.Channel()
	if nil != err {
		panic(err)
	}

	queue, err := ch.QueueDeclare("hey", true, false, false, false, nil)
	if nil != err {
		panic(err)
	}

	//设置成confirm模式
	if err = ch.Confirm(false); nil != err {
		panic(err)
	}

	// 判断成功或失败将消息发送给exchange
	// 猜测当没有给exchange发送成功的时候,这个库会自动重传,因为测试的时候,
	// 这里的confirms channel一直都在返回
	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	go func() {
		for {
			select {
			case v := <-confirms:
				//成功返回true, 失败返回false
				fmt.Println(v.Ack, v.DeliveryTag)
			default:

			}
		}
	}()

	// 判断是否成功由exchange到queue
	res := ch.NotifyReturn(make(chan amqp.Return))
	go func() {
		for {
			select {
			// 当失败才有返回,若成功进入queue则不会有channel返回
			case r := <-res:
				if r.ReplyCode != 0 {
					fmt.Println(r)
				}
			default:
			}
		}
	}()

	body := "hello"
	for i := 0; i < 1; i++ {
		err := ch.Publish("", queue.Name, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		//猜测这个地方publish是异步的,这里的err可能就是单纯的是ch链接断了发失败了.
		if nil != err {
			panic("消息发送的时候挂了:" + err.Error())
		}
	}
}
