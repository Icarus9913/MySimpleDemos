package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//声明算数运算结构体
type Arith struct {}

//声明接受的参数结构体
type ArithRequest struct {
	A,B int
}

//声明返回客户端参数结构体
type ArithResponse struct {
	Pro int		//乘积
	Quo int		//商
	Rem int		//余数
}

//乘法运算
func (this *Arith)Multiply(req ArithRequest, res *ArithResponse)error  {
	res.Pro = req.A * req.B
	return nil
}

//商和余数
func (this * Arith)Divide( req ArithRequest,res *ArithResponse)error  {
	if req.B==0{
		return errors.New("除数不能为0")
	}
	//商
	res.Quo = req.A / req.B
	//余数
	res.Rem = req.A % req.B
	return nil
}

//普通方式
/*func main()  {
	//注册服务
	rpc.Register(new(Arith))
	//采用http作为rpc载体
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8081", nil)
	if nil!=err{
		log.Fatal(err)
	}
}*/

//jsonRPC编码方式
func main()  {
	//注册服务
	rpc.Register(new(Arith))
	//监听服务
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if nil!=err{
		log.Fatal(err)
	}
	//循环监听服务
	for {
		conn, err:=lis.Accept()
		if nil!=err{
			continue
		}
		//协程
		go func(conn net.Conn) {
			fmt.Println("new a Client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}


