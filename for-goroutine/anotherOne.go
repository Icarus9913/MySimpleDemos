package main

import (
	"fmt"
	"time"
)

/*
注意1:
	变量名是用来引用变量所存储数据的标识符(Identifier),它实际上代表变量在内存中的地址,可以使用取址符&获得
	从变量中取值,就是通过变量名找到相应的内存地址,再从该存储单元中读取数据
注意2:
	非指针数据类型的值,也能调用其指针方法,这是因为Go在内部做了自动转换.
	例如,若add方法是指针方法,那么表达式 a1.add(2)会被自动转换为(&a1).add(2)
注意3:
	值类型实参   调用 值类型方法
	指针类型实参 调用 值类型方法(自动解引用)
	值类型      调用 指针类型方法(自动取引用)
	指针类型    调用 指针类型方法
*/

type myString string

func main() {
	//values := []myString{"ni", "hao", "a", "da", "ge"}
	//bad1(values)
	goodMethod2()
	time.Sleep(1 * time.Second)
}

/*
	每次for循环中v的地址是相同的,复用的.所以打印这个值的时候是根据地址打印.
	而主线程运行的比goroutine快,所以在goroutine里打印的时候,主线程中已经把v的值改变了
*/
func bad1(values []myString) {
	for _, v := range values {
		go func() {
			fmt.Println(v)
		}()
	}
}

/*
	每次for循环中都会声明一个临时变量去存v,而每次打印的时候,打印的val的值不一样.
	即启动goroutine之前,val的值已经确定了
*/
func good1(values []myString) {
	for _, v := range values {
		val := v
		go func() {
			fmt.Println(val)
		}()
	}
}

//每次for循环的时候,已经把v交给匿名函数的行参了.因此每次print的时候,已经确定好了val的值
func good2(values []myString) {
	for _, v := range values {
		go func(val myString) {
			fmt.Println("打印地址:",&val)
			fmt.Println(val)
		}(v)
	}
}

//解释同bad1
func badMethod(values []myString) {
	for _, v := range values {
		go v.myMethod()
	}
}

//解释同good1
func goodMethod1(values []myString) {
	for _, v := range values {
		newVal := v
		go newVal.myMethod()
	}
}

/*
	v是指针类型,调用对应的方法,其方法接收者是指针类型.调用的方法是v对应其指针的方法,
	而每次for循环的时候,这个指针都是不一样的,所以OK
*/
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
