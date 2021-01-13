package main

import (
	"Rpc_Master/RpcAndProto/message"
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

//订单服务
type OrderService struct {
}

func (os *OrderService)GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"2020100500001":message.OrderInfo{OrderId: "2020100500001",OrderName: "衣服",OrderStatus: "已付款"},
		"2020100510001":message.OrderInfo{OrderId: "2020100510001",OrderName: "零食",OrderStatus: "已付款"},
		"2020100510002":message.OrderInfo{OrderId: "2020100510002",OrderName: "手机",OrderStatus: "未付款"},
	}

	current := time.Now().Unix()
	if request.TimeStamp > current{
		*response = message.OrderInfo{OrderId: "0",OrderName: "",OrderStatus: "订单信息异常"}
	}else {
		result := orderMap[request.OrderId]
		if result.OrderId != ""{
			*response = orderMap[request.OrderId]
		}else {
			return errors.New("server error")
		}
	}
	return nil
}

func main()  {
	orderService := new(OrderService)
	rpc.Register(orderService)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8081")
	if nil!=err{
		panic(err.Error())
	}
	http.Serve(listen,nil)

}