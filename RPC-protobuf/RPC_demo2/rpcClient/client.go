package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//声明接受的参数结构体
type ArithRequest struct {
	A,B int
}

//声明返回客户端参数结构体
type ArithResponse struct {
	//乘积
	Pro int
	//商
	Quo int
	//余数
	Rem int
}

//调用服务
func main()  {
	//连接远程rpc
	//普通方式
	//rpcclient, err := rpc.DialHTTP("tcp", "localhost:8081")

	//jsonrpc方式
	rpcclient, err := jsonrpc.Dial("tcp", "localhost:8081")

	if nil!=err{
		log.Fatal(err)
	}
	req := ArithRequest{9,2}
	var res ArithResponse
	//调用乘积
	err = rpcclient.Call("Arith.Multiply", req, &res)
	if nil!=err{
		log.Fatal(err)
	}
	fmt.Printf("%d*%d=%d\n",req.A,req.B,res.Pro)
	//调用商
	err = rpcclient.Call("Arith.Divide", req, &res)
	if nil!=err{
		log.Fatal(err)
	}
	fmt.Printf("%d/%d,商=%d, 余数=%d\n",req.A,req.B,res.Quo,res.Rem)
}