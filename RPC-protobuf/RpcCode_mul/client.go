package main

import (
	"Rpc_Master/RpcCode_mul/param"
	"fmt"
	"net/rpc"
)

func main()  {
	client, err := rpc.DialHTTP("tcp", "localhost:8082")
	if nil!=err{
		panic(err.Error())
	}

	var result *float32
	addParam := &param.AddParam{Args1: 1.2,Args2: 2.3}
	err = client.Call("MathUtil.Add", addParam, &result)
	if nil!=err{
		panic(err.Error())
	}
	fmt.Println("计算结果:", *result)

}
