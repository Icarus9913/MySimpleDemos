package main

import (
	"Rpc_Master/6ServerAndClientStreamCode/message"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
)

//订单服务实现
type OrderServiceImpl struct {
}

func (os *OrderServiceImpl)GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error {
	for{
		orderRequest, err := stream.Recv()
		if err==io.EOF{
			fmt.Println("读取数据结束")
			return err
		}
		if nil!=err{
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]message.OrderInfo{
			"2020100500001":message.OrderInfo{OrderId: "2020100500001",OrderName: "衣服",OrderStatus: "已付款"},
			"2020100510001":message.OrderInfo{OrderId: "2020100510001",OrderName: "零食",OrderStatus: "已付款"},
			"2020100510002":message.OrderInfo{OrderId: "2020100510002",OrderName: "手机",OrderStatus: "未付款"},
		}
		result:=orderMap[orderRequest.GetOrderId()]
		//发送数据
		err=stream.Send(&result)
		if err==io.EOF{
			fmt.Println(err)
			return err
		}
		if nil!=err{
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func main()  {
	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server,new(OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8091")
	if nil!=err{
		panic(err.Error())
	}
	server.Serve(listen)
}

