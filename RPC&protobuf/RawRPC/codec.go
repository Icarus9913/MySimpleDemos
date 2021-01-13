package RawRPC

import (
	"bytes"
	"encoding/gob"
)

//定义数据格式和遍解码

//定义RPC交互的数据格式
type RPCData struct {
	//访问的函数
	Name string
	//访问时传的参数
	Args []interface{}
}

//编码
func encode(data RPCData)([]byte,error)  {
	var buf bytes.Buffer
	//得到字节数组的编码器
	bufEnc := gob.NewEncoder(&buf)
	//对数据编码
	if err := bufEnc.Encode(data);err!=nil{
		return nil, err
	}
	return buf.Bytes(),nil
}

//解码
func decode(b []byte)(RPCData,error)  {
	buf := bytes.NewBuffer(b)
	//返回字节数组解码器
	bufDec:=gob.NewDecoder(buf)
	var data RPCData
	//对数据解码
	if err:=bufDec.Decode(&data); err!=nil{
		return data,err
	}
	return data,nil
}