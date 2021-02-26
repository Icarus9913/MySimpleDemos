package main

import (
	"Rpc_Master/6ServerStreamCode/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
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
	request := message.OrderRequest{TimeStamp: time.Now().Unix()}
	orderInfosClient, err := orderServiceClient.GetOrderInfos(context.Background(), &request)

	for{
		orderInfo, err := orderInfosClient.Recv()
		if err==io.EOF{		//end of file
			fmt.Println("读取结束")
			return
		}
		if nil!=err{
			panic(err.Error())
		}
		fmt.Println("读取到的信息：",orderInfo)
	}

	
}
