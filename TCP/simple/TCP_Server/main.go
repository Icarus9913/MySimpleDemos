package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	service := ":5000"
	tcpAddr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	for{
		conn,err := listener.Accept()
		if nil!=err{
			continue
		}
		go handleClient(conn)
		//conn.Close()
	}
}

func handleClient(conn net.Conn)  {
	defer conn.Close()
	var buf [512]byte
	for  {
		n,err := conn.Read(buf[0:])
		if nil!=err{
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client",rAddr.String(),string(buf[0:n]))
		_,err2:=conn.Write([]byte("\nWelcome client!"))
		if nil!=err2{
			return
		}
	}
}

func checkError(err error)  {
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}
