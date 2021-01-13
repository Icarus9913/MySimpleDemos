package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

//测试读写
func TestSession_ReadWrite(t *testing.T) {
	//定义监听IP和端口
	addr:="127.0.0.1:8000"
	//定义传输的数据
	my_data:="hello"
	//等待组
	wg:=sync.WaitGroup{}
	//协程1个读，1个写
	wg.Add(2)

	//写数据协程
	go func() {
		defer wg.Done()
		//创建tcp连接
		lis, err := net.Listen("tcp", addr)
		if nil!=err{
			t.Fatal(err)
		}
		conn, _ := lis.Accept()
		s := Session{conn: conn}
		//写数据
		err = s.Write([]byte(my_data))
		if nil!=err{
			t.Fatal(err)
		}
	}()

	time.Sleep(1*time.Second)
	//读数据协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if nil!=err{
			t.Fatal(err)
		}
		s:=Session{conn: conn}
		//读数据
		data,err:=s.Read()
		if nil!=err{
			t.Fatal(err)
		}
		if string(data)!=my_data{
			t.Fatal(err)
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}
