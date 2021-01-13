package main

import (
	"Rpc_Master/6ClientStreamCode/message"
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
	orderServiceClient := message.NewOrderServiceClient(conn)

	fmt.Println("客户端请求RPC调用：客户端流模式")
	orderMap := map[string]message.OrderInfo{
		"2020100500001":message.OrderInfo{OrderId: "2020100500001",OrderName: "衣服",OrderStatus: "已付款"},
		"2020100510001":message.OrderInfo{OrderId: "2020100510001",OrderName: "零食",OrderStatus: "已付款"},
		"2020100510002":message.OrderInfo{OrderId: "2020100510002",OrderName: "手机",OrderStatus: "未付款"},
	}
	//调用服务方法
	addOrderListClient, err := orderServiceClient.AddOrderList(context.Background())
	if nil!=err{
		panic(err.Error())
	}
	//调用方法发送流数据
	for _,info:=range orderMap{
		err = addOrderListClient.Send(&info)
		if nil!=err{
			panic(err.Error())
		}
	}
	for{
		orderInfo, err := addOrderListClient.CloseAndRecv()
		if err==io.EOF{
			fmt.Println("读取数据结束")
			return
		}
		if nil!=err{
			panic(err.Error())
		}
		fmt.Println(orderInfo.GetOrderStatus())
	}
}