package main

type S struct {}

func f(x interface{})  {}

func g(x *interface{})  {}


func main()  {
	s := S{}		//type:S
	p := &s			//type:*S

	//f(x interface{})参数类型为interface可以接收 S类型 和 *S类型
	f(s)
	f(p)

	//g(x interface{})参数类型为*interface不能接收 S类型 和 *S类型

/*	g(s)
	g(p)
*/

	
}