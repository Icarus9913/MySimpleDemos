package main

import (
	"Rpc_Master/7grpcSSLcode/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main()  {
	//TSL连接
	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "go-grpc-example") //common name
	if nil!=err{
		panic(err)
	}
	grpc.WithInsecure()

	//Dial连接
	conn, err := grpc.Dial("localhost:8092", grpc.WithTransportCredentials(creds))
	if nil!=err{
		panic(err)
	}
	defer conn.Close()

	serviceClient := message.NewMathServiceClient(conn)

	addArgs := message.RequestArgs{Args1: 3, Args2: 5}

	response, err := serviceClient.AddMethod(context.Background(), &addArgs)
	if err != nil {
		grpclog.Fatal(err.Error())
	}

	fmt.Println(response.GetCode(), response.GetMessage())
}
