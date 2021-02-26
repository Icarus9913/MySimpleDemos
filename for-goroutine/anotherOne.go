package main

import (
	"fmt"
	"time"
)

type myString string

func main() {
	values := []myString{"ni", "hao", "a", "da", "ge"}
	bad1(values)
	time.Sleep(1 * time.Second)
}

//理解为启动协程所耗时间比for循环结束时间要长,所以都是打印出"ge"
func bad1(values []myString) {
	for _, v := range values {
		go func() {
			fmt.Println(v)
		}()
	}
}

//理解为启动goroutine之前,val的值已经确定了,不会存在goroutine之间的竞争
func good1(values []myString) {
	for _, v := range values {
		val := v
		go func() {
			fmt.Println(val)
		}()
	}
}

//理解为print的时候,已经确定好了val的值
func good2(values []myString) {
	for _, v := range values {
		go func(val myString) {
			fmt.Println(val)
		}(v)
	}
}

func badMethod(values []myString) {
	for _, v := range values {
		go v.myMethod()
	}
}

func goodMethod1(values []myString) {
	for _, v := range values {
		newVal := v
		go newVal.myMethod()
	}
}

func goodMethod2() {
	ni := myString("ni")
	hao := myString("hao")
	a := myString("a")
	da := myString("da")
	ge := myString("ge")
	values := []*myString{&ni, &hao, &a, &da, &ge}
	for _, v := range values {
		go v.myMethod()
	}
}

func (v *myString) myMethod() {
	fmt.Println(*v)
}
