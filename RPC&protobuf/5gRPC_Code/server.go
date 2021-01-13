package main

import (
	"Rpc_Master/5gRPC_Code/message"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

type orderServiceImpl struct {}

//具体的方法实现
func (os *orderServiceImpl)GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error){
	orderMap := map[string]message.OrderInfo{
		"2020100500001":message.OrderInfo{OrderId: "2020100500001",OrderName: "衣服",OrderStatus: "已付款"},
		"2020100510001":message.OrderInfo{OrderId: "2020100510001",OrderName: "零食",OrderStatus: "已付款"},
		"2020100510002":message.OrderInfo{OrderId: "2020100510002",OrderName: "手机",OrderStatus: "未付款"},
	}
	var response *message.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current{
		*response = message.OrderInfo{OrderId: "0",OrderName: "",OrderStatus: "订单信息异常"}
	}else {
		result := orderMap[request.OrderId]
		if result.OrderId != ""{
			fmt.Println(result)
			return &result,nil
		}else {
			return nil,errors.New("server error")
		}
	}
	return response,nil
}

func main()  {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server,new(orderServiceImpl))
	listen, err := net.Listen("tcp", ":8090")
	if nil!=err{
		panic(err)
	}
	fmt.Println("server running...")
	server.Serve(listen)

}