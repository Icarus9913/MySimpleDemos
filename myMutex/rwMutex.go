package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}


var Password = secret{password: "myPassword"}

func Change(c *secret, pass string) {
	c.RWM.Lock()
		fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	//fmt.Println("打印C.PASSWORD:",c.password)
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
		fmt.Println("show")
	time.Sleep(3*time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

//排他索表明只有一个showWithLock()函数能读取secret结构体的password字段
func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3*time.Second)
	defer c.M.Unlock()
	return c.password
}

func main()  {
	var showFunction = func(c *secret) string {return ""}
	if len(os.Args)!=2{
		fmt.Println("Using sync.RWMutex!")
		showFunction = show				//无参数时RWMutex
	}else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock		//有参数时Mutex
	}

	var waitGroup sync.WaitGroup
	fmt.Println("Pass:", showFunction(&Password),"\n")

	for i:=0;i<1;i++{
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:",showFunction(&Password))
		}()

		go func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			Change(&Password,"123456")
		}()
		waitGroup.Wait()
		fmt.Println("最后打印:",showFunction(&Password))
	}
}