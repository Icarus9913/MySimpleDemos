package RawRPC

import (
	"net"
	"reflect"
)

//声明客户端
type Client struct {
	conn net.Conn
}

//创建客户端对象
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

//实现通用的RPC客户端
//绑定RPC访问的方法
//传入访问的函数名

//函数具体实现在Server端，Client只有函数原型
//使用MakeFunc()完成原型到函数的调用

//fPtr指向函数原型
//例如xxx.calRPC("queryUser",&query)
func (c *Client)callRPC(rpcName string, fPtr interface{})  {
	//通过反射，获取fPtr未初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	 //另一个函数，作用是对地一个函数参数操作
	 //完成与Server的交互
	 f:= func(args []reflect.Value)[]reflect.Value {
	 	//处理输入的参数
	 	inArgs := make([]interface{},0,len(args))
	 	for _,args:=range args{
	 		inArgs=append(inArgs,args.Interface())
		}
		//创建连接
		cliSession:= NewSession(c.conn)
		//编码数据
		reqRPC:= RPCData{Name: rpcName,Args: inArgs}
		b,err:= encode(reqRPC)
		if nil!=err{
			panic(err)
		}
		//写出数据
		 err = cliSession.Write(b)
		 if nil!=err{
		 	panic(err)
		 }
		 //读取相应数据
		 respBytes,err:=cliSession.Read()
		 if nil!=err{
		 	panic(err)
		 }
		 //解码数据
		 respRPC, err := decode(respBytes)
		 if nil!=err{
		 	panic(err)
		 }
		 //处理服务端返回的数据
		 outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		 for i,arg := range respRPC.Args{
		 	//必须进行 nil 转换
		 	if arg==nil{
		 		//必须填充一个真正的类型，不能是nil
		 		outArgs=append(outArgs,reflect.Zero(fn.Type().Out(i)))
		 		continue
			}
			outArgs=append(outArgs,reflect.ValueOf(arg))
		 }
		 return outArgs
	 }



	//参数1：一个未初始化函数的方法值，类型是reflect.Type
	//参数2：另一个函数，作用是对第一个函数参数操作
	//返回reflect.Value类型
	//MakeFunc使用传入的函数原型，创建一个绑定 参数2 的新函数
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)

}
