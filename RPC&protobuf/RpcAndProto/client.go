package main

import (
	"Rpc_Master/RpcAndProto/message"
	"fmt"
	"net/rpc"
	"time"
)

func main()  {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if nil!=err{
		panic(err)
	}
	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "2020100510002", TimeStamp: timeStamp}

	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo",request,&response)
	if nil!=err{
		panic(err)
	}
	fmt.Println(*response)

}

