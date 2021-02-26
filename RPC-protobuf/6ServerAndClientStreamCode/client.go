package main

import (
	"Rpc_Master/6ServerAndClientStreamCode/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
)

func main()  {
	//1.Dial连接
	conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure())
	if nil!=err{
		panic(err)
	}
	defer conn.Close()

	fmt.Println("客户端请求RPC调用：双向流模式")
	orderIds := []string{"2020100500001","2020100510001","2020100510002"}

	orderServiceClient := message.NewOrderServiceClient(conn)
	orderInfosClient, err := orderServiceClient.GetOrderInfos(context.Background())
	for _,orderId := range orderIds{
		orderRequest:=message.OrderRequest{OrderId: orderId}
		err := orderInfosClient.Send(&orderRequest)
		if nil!=err{
			panic(err)
		}
	}

	//关闭
	orderInfosClient.CloseSend()
	for{
		orderInfo, err := orderInfosClient.Recv()
		if err==io.EOF{
			fmt.Println("读取结束")
			return
		}
		if nil!=err{
			panic(err)
		}
		fmt.Println("读取到的信息：",orderInfo)
	}
}
