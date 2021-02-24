package main

import (
	"fmt"
	"reflect"
)

type Userer interface {
	Say()
}

type User struct {
	MyName string
	Say    func() //此处的类型无所谓,因为后面只需要此处的名称"Say"这个字符串
}

type Person struct{}

func (p Person) Say() {
	fmt.Println("打印")
}

func main() {
	var u User
	valueOf := reflect.ValueOf(u)
	field := valueOf.Type().Field(1) //这里是User结构体中Say这个字段的所有信息,,,而我们只需要他的名字字符串"Say"

	var us Userer = Person{}
	sValue := reflect.ValueOf(us)

	myFunc := sValue.MethodByName(field.Name) //对象根据字符串名去调用对应的方法

	args := []reflect.Value{} //这是那个调用函数的参数,因为这里Say()方法中没有参数,所以不需要赋值
	myFunc.Call(args)         //call是该方法的调用
}
