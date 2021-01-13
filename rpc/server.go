package rpc

import (
	"fmt"
	"net"
	"reflect"
)

//声明服务端
type Server struct {
	addr string							//地址
	funcs map[string]reflect.Value		//服务端维护的函数名到函数反射值的map
}

//创建服务端对象
func NewServer(addr string) *Server  {
	return &Server{addr: addr,funcs: make(map[string]reflect.Value)}
}


//服务端绑定注册方法
//将函数名与函数真正实现对应起来
//第一个参数为函数名，第二个传入真正的函数
func (s *Server)Register(rpcName string,f interface{})  {
	if _,ok := s.funcs[rpcName];ok{
		return
	}
	//map中没有值，则将映射添加进map，便于调用
	fVal:=reflect.ValueOf(f)
	s.funcs[rpcName]=fVal
}

//服务端等待调用
func (s *Server)Run()  {
	//监听
	lis, err := net.Listen("tcp", s.addr)
	if nil!=err{
		fmt.Printf("监听 %s err:%v",s.addr,err)
	}
	for{
		//拿到连接
		conn,err:=lis.Accept()
		if nil!=err{
			fmt.Printf("accept err:=%v",err)
			return
		}
		//创建会话
		srvSession := NewSession(conn)
		//RPC读取数据
		b, err := srvSession.Read()
		if nil!=err{
			fmt.Printf("read err:%v",err)
			return
		}
		//对数据解码
		rpcData, err := decode(b)
		if nil!=err{
			fmt.Printf("read err:%v",err)
			return
		}
		//根据读取到的数据的Name，得到调用的函数名
		f,ok := s.funcs[rpcData.Name]
		if !ok{
			fmt.Printf("函数%s不存在",rpcData.Name)
			return
		}
		//解析遍历客户端传来的参数，放到一个数组中
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _,arg:=range rpcData.Args{
			inArgs=append(inArgs,reflect.ValueOf(arg))
		}
		//反射调用方法，传入参数
		out := f.Call(inArgs)
		//解析遍历执行结果，放到一个数组中
		outArgs := make([]interface{},0,len(out))
		for _,o := range out{
			outArgs=append(outArgs,o.Interface())
		}
		//包装数据，返回给客户端
		respRPCData:=RPCData{rpcData.Name,outArgs}
		//编码
		respBytes,err:=encode(respRPCData)
		if nil!=err{
			fmt.Printf("encode err:%v\n",err)
			return
		}
		//使用RPC写出数据
		err = srvSession.Write(respBytes)
		if nil!=err{
			fmt.Printf("session write err:%v\n",err)
			return
		}
	}
}