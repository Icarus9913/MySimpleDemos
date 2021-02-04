package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

//set函数的目的是共享变量的值
func set(newValue int) {
	writeValue <- newValue
}

//read1函数的目的是读取已保存变量的值
func read1() int {
	return <-readValue
}

/*
	当有读请求时,read1()函数从由monitor()函数控制的readValue管道读取数据.返回保存在value变量中的当前值
	若想修改存储值,就是调用set函数,但是从wqriteValue里拿数据也是由select语句处理
	结果就是,不使用monitor()函数,没人可以处理value共享变量
*/
func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d, ", value)
		case readValue <-value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random number.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var w sync.WaitGroup
	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}
	w.Wait()

	fmt.Printf("\nLast value: %d\n", read1())
}
