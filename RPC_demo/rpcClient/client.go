package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width,Height int
}

//调用服务
func main()  {
	//1、连接远程rpc服务
	rpcclient, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if nil!=err{
		log.Fatal(err)
	}
	//2、调用远程方法
	//定义接受服务端传回来的计算结果的变量
	result := 0
	//求面积
	err = rpcclient.Call("Rect.Area", Params{50, 100}, &result)
	if nil!=err{
		log.Fatal(err)
	}
	fmt.Println("面积：",result)

	//求周长
	err = rpcclient.Call("Rect.Perimeter", Params{50, 100}, &result)
	if nil!=err{
		log.Fatal(err)
	}
	fmt.Println("周长：",result)

}