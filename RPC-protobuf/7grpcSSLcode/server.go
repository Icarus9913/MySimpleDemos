package main

import (
	"Rpc_Master/7grpcSSLcode/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net"
)

type MathManager struct {
}

func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
	fmt.Println(" 服务端 Add方法 ")
	result := request.Args1 + request.Args2
	fmt.Println(" 计算结果是：", result)
	response = new(message.Response)
	response.Code = 1;
	response.Message = "执行成功"
	return response, nil
}

func main()  {
	//TSL认证
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if nil!=err{
		grpclog.Fatal("加载证书文件失败",err)
	}
	server := grpc.NewServer(grpc.Creds(creds))

	listen, err := net.Listen("tcp", ":8092")
	if nil!=err{
		panic(err)
	}
	server.Serve(listen)

}


