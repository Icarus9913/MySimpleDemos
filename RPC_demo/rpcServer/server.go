package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//服务端，求矩形面积和周长

//声明矩形对象
type Rect struct {}

//声明参数结构体，字段首字母大写
type Params struct {
	Width,Height int		//长和宽
}

//定义求矩形面积的方法
func (r *Rect)Area(p Params,result *int) error {
	*result = p.Width * p.Height
	return nil
}

//定义求矩形周长的方法
func (r *Rect)Perimeter(p Params, result *int) error  {
	*result = (p.Width + p.Height) * 2
	return nil
}

func main()  {
	//1、注册服务
	rect := new(Rect)
	rpc.Register(rect)

	//2、把服务处理绑定到http协议上
	rpc.HandleHTTP()

	//3、监听服务，等待客户端调用面积和周长的方法
	err := http.ListenAndServe(":8080",nil)
	if nil!=err{
		log.Fatal(err)
	}
}