package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if nil!=err{
		panic(err.Error())
	}

	var req float32		//请求值
	req = 10

/*	//同步的调用方式
	var resp *float32	//返回值
	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	if nil!=err{
		panic(err.Error())
	}
	fmt.Println(*resp)
*/
	//异步的调用方式
	var respSync *float32
	syncCall := client.Go("MathUtil.CalculateCircleArea",req,&respSync,nil)
	replyDone := <-syncCall.Done
	fmt.Println(replyDone)
	fmt.Println(*respSync)		//只有syncCall发起请求并接收数据后，这里respSync才有数据


}
