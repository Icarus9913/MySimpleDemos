package main

import (
	"sync/atomic"
	"time"
)
/*
	自旋锁是指当一个线程在获取锁的时候，如果锁已经被其他线程获取，则该线程将循环等待，
	然后不断的判断是否能够被成功获取，直到获取了锁才会退出循环.
	获取锁的线程一直处于活跃状态，但是并没有执行任何有效的任务，使用这种锁会造成busy-waiting。
*/

type Locker interface {
	Lock()
	Unlock()
}

type Spin int32

/*
	func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
	CompareAndSwapInt32原子性的比较*addr和old，如果相同则将new赋值给*addr并返回真。
*/
func (s *Spin)Lock()  {
	for !atomic.CompareAndSwapInt32((*int32)(s),0,1){}	//s等于0则s变成1，返回true，非true不执行
}																	//s等于1，则返回false，非false则一直执行

/*
	func StoreInt32(addr *int32, val int32)
	StoreInt32原子性的将val的值保存到*addr。
*/
func (s *Spin)Unlock()  {
	atomic.StoreInt32((*int32)(s),0)
}

func main() {
	var l Locker = new(Spin)

	var n int
	for i := 0;i<2;i++{
		go routine(i,&n,l,500*time.Millisecond)
	}
	select {}
}

func routine(i int,v *int,l Locker,d time.Duration)  {
	for {
		func(){
			l.Lock()
			defer l.Unlock()
			*v++
			println(*v,i)
			time.Sleep(d)
		}()
	}
}