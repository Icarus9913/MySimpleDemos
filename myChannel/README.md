## Let's talk about some shits with Channel!

如果chan没有缓冲,所以发送的地方只有在有目标接收的时候才能发出

注意:
如果chan不**close**的话!!!!!,则会在以下两种情况下出现{死锁}.  
1.因为for range无限迭代下去;   
2.对于for循环,channel的源码下,close后会对已经close的chan变量,发送0,false
>for range或者  
>```
>for{
>  case v,ok := <-channel 
> }
> ```  

--------------------------------------------------  
--------------------------------------------------  
  
## 注意事项2:
- 对于带有缓冲的的chan变量,如果塞了值之后close掉,再读可以读到数据,读完之后就都是0
- 对于无缓冲的chan变量,如果塞了值后close掉,读不到之前的数据,读出来的都是0
>   测试一个channel是否被关闭: 使用 val,ok := <- c    
>   第二个结果是一个布尔值ok, true表示从channels接收到值, false表示channels已经被关闭并且里面没有值可以接收

```go
package main

import (
	"fmt"
	"time"
)

/*
    带有缓冲的变量c,给他发送值后close,还能读到值,注意此时的ok是true!!!!
*/
func have() {
	c := make(chan int,1)   //有缓冲
	
	go func(c chan int) {
		c <- 22
	}(c)
	time.Sleep(100 * time.Millisecond)
	close(c)

	val, ok := <-c
	fmt.Println(ok, ":", val)   //true : 22
}

/*
    无缓冲的变量c,发送值后close,无法读到值,注意此时的ok是false!!!
*/
func notHave()  {
	c := make(chan int) //无缓冲
	
	go func(c chan int) {
		c <- 22
	}(c)
	time.Sleep(100 * time.Millisecond)
	close(c)
	
	val, ok := <-c
	fmt.Println(ok, ":", val)   //false : 0
}   
```

