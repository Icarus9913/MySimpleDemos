package main

import (
	"fmt"
	"reflect"
)

type Actor struct {}

type User struct {
	ID int
	Name string
}

type Student struct {
	Age int
}

type invokeFunc func(u User)

func (a Actor)PrintID(u User) int {
	fmt.Println("打印user的ID,",u.ID)
	return u.ID
}

func (a Actor)PrintName(u User) string {
	fmt.Println("打印user的Name,",u.Name)
	return u.Name
}




func (a Actor)Exports() []interface{} {
	return []interface{}{
		1: a.PrintID,
		2: a.PrintName,
	}
}


func main()  {
	/*	user := User{
		ID: 1,
		Name: "小李",
	}
	value := reflect.ValueOf(user)
	check := value.MethodByName("PrintID")
	args := make([]reflect.Value,0)
	check.Call(args)*/

	var a Actor
	exports := a.Exports()
	for _, m := range exports{
		//忽略掉数组0,因为数组0对应的元素为空
		if nil==m{
			continue
		}

		meth := reflect.ValueOf(m)
		_ = meth.Type()		//输出对应的类型及返回值,例如:func(main.User) string
		_ = reflect.TypeOf(m)		///输出对应的类型及返回值,例如:func(main.User) int



		//u := param.Interface().(*int)

		meth.Call([]reflect.Value{
			reflect.ValueOf(User{
				2,
				"abc",
			}),

		})

	}
}
