package main

import (
	"Rpc_Master/RpcCode_mul/param"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//数学计算
type MathUtil struct {}

//该方法向外暴露：提供计算圆形面积的服务
func (mu *MathUtil)Add(param param.AddParam, resp *float32) error  {
	*resp = param.Args1 + param.Args2 //实现两数相加的功能
	return nil
}

func main()  {
	//1.初始化指针数据类型
	mathUtil := new(MathUtil)	//初始化指针数据类型
	//2.调用net/rpc包的功能将服务对象进行注册
	err := rpc.RegisterName("MathUtil",mathUtil)			//注意与RpcCode中的改变
	if nil!=err{
		panic(err.Error())
	}
	//3.通过该函数把mathUtil中提供的服务注册到HTTP协议上 方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()
	//4.在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8082")
	if nil!=err{
		panic(err.Error())
	}
	fmt.Println("Server start...")
	http.Serve(listen,nil)
}
