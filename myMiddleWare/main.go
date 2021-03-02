package main

import "fmt"

type HandleFunc func(request *Request)

type Request struct {
	index       int
	middlewares []HandleFunc //中间件+处理函数
}

func main() {
	//假设来了一个请求
	request := NewRequest()

	//注册了一些中间件
	request.RegisterMiddlewares(logger,recovery, func(request *Request) {
		//业务处理函数作为中间件的最后一环
		fmt.Println("我是业务逻辑")
	})

	//然后开始请求
	request.Next()
}

//生成请求
func NewRequest() (request *Request) {
	request = &Request{
		index:       0,
		middlewares: make([]HandleFunc, 0),
	}
	return
}

//注册中间件
func (request *Request) RegisterMiddlewares(middlewares ...HandleFunc) {
	for _, mid := range middlewares {
		request.middlewares = append(request.middlewares, mid)
	}
}

//执行中间件
func (request *Request) Next() {
	index := request.index
	if index >= len(request.middlewares) {
		return
	}

	request.index++
	request.middlewares[index](request)
}

//日志中间件
func logger(request *Request)  {
	fmt.Println("请求开始")  //before 下一个中间件
	request.Next()
	fmt.Println("请求结束")  //after  下一个中间件
}

//panic捕获中间件
func recovery(request *Request)  {
	defer func() {
		recover()
		fmt.Println("我确保panic被捕获")
	}()
	request.Next()
}


