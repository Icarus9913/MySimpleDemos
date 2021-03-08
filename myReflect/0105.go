package main

import (
	"fmt"
	"reflect"
)

type Actor struct{}

type Account struct {
	ID   int
	Name string
}

type Student struct {
	Age int
}

type invokeFunc func(ac Account)

func (a Actor) PrintID(ac Account) int {
	fmt.Println("打印account的ID,", ac.ID)
	return ac.ID
}

func (a Actor) PrintName(ac Account) string {
	fmt.Println("打印account的Name,", ac.Name)
	return ac.Name
}

func (a Actor) Exports() []interface{} {
	return []interface{}{
		1: a.PrintID,
		2: a.PrintName,
	}
}

func main() {
	/*	account := Account{
			ID: 1,
			Name: "小李",
		}
		value := reflect.ValueOf(account)
		check := value.MethodByName("PrintID")
		args := make([]reflect.Value,0)
		check.Call(args)*/

	var a Actor
	exports := a.Exports()
	for _, m := range exports {
		//忽略掉数组0,因为数组0对应的元素为空
		if nil == m {
			continue
		}

		meth := reflect.ValueOf(m)
		_ = meth.Type()       //输出对应的类型及返回值,例如:func(main.Account) string
		_ = reflect.TypeOf(m) //输出对应的类型及返回值,例如:func(main.Account) int

		//ac := param.Interface().(*int)

		meth.Call([]reflect.Value{
			reflect.ValueOf(Account{
				2,
				"abc",
			}),
		})
	}
}
