package main

import (
	"Rpc_Master/5gRPC_Code/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main()  {
	//1.Dial连接
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if nil!=err{
		panic(err)
	}
	defer conn.Close()
	orderServiceClient := message.NewOrderServiceClient(conn)
	orderRequest := &message.OrderRequest{OrderId: "2020100500001", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if nil!=orderInfo{
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}else {
		fmt.Println(err.Error())
	}
}