package main

import (
	"Rpc_Master/6ServerStreamCode/message"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

//订单服务实现
type OrderServiceImpl struct {
}

//获取订单信息
func (os *OrderServiceImpl)	GetOrderInfos(request *message.OrderRequest, stream message.OrderService_GetOrderInfosServer) error {
	fmt.Println("服务端流RPC模式")

	orderMap := map[string]message.OrderInfo{
		"2020100500001":message.OrderInfo{OrderId: "2020100500001",OrderName: "衣服",OrderStatus: "已付款"},
		"2020100510001":message.OrderInfo{OrderId: "2020100510001",OrderName: "零食",OrderStatus: "已付款"},
		"2020100510002":message.OrderInfo{OrderId: "2020100510002",OrderName: "手机",OrderStatus: "未付款"},
	}
	for id,info := range orderMap{
		if time.Now().Unix()>=request.TimeStamp{
			fmt.Println("订单序列号ID：",id)
			fmt.Println("订单详情：",info)
			//通过流模式发送给客户端
			stream.Send(&info)
		}
	}
	return nil
}

func main()  {
	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server,new(OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8090")
	if nil!=err{
		panic(err.Error())
	}
	server.Serve(listen)

}