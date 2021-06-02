package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	该协程池的设计增加了release操作.
	一个pool中包含有一个EntryJob的channel以及一个worker的切片.
	一个worker中包含有这个pool的实例,为了从这个pool的EntryChannel中拿到数据.
	另外每个worker中含有自己的jobChannel.

	release操作就是先close掉通知channel,然后再向每个worker的jobChannel放入一个nil任务,
	当worker判断接收到的任务是nil时候,停止执行.再将每个worker置空,接着pool的worker切片置空即可.
	随后gc会将未引用的worker channel变量给清除.
*/

var stopCh chan struct{}

//定义任务task类型,每一个任务task都可以抽象成一个函数
type task struct {
	f func() error
}

//通过newTask来创建一个task
func newTask(f func() error) *task {
	t := task{
		f: f,
	}
	return &t
}

//执行Task任务的方法
func (t *task) Execute() {
	t.f() //调用任务所绑定的函数s
}

/*有关协程池的定义及操作*/
//定义池类型
type Pool struct {
	//对外接收Task的入口
	EntryChannel chan *task
	workers      []*Worker
}

type Worker struct {
	pool *Pool

	//协程池内部的任务就绪队列
	JobsChannel chan *task
}

//创建一个协程池
func newPool(cap int) *Pool {
	p := &Pool{}

	w := func() []*Worker {
		t := make([]*Worker, cap)
		for i := 0; i < cap; i++ {
			t[i] = &Worker{
				pool:        p,
				JobsChannel: make(chan *task, 10),
			}
		}
		return t
	}
	p.workers = w()
	p.EntryChannel = make(chan *task, len(p.workers)*10)
	return p
}

func (p *Pool) run() {
	for _, w := range p.workers {
		go w.run()
	}
}

func (w *Worker) run() {
	go func() {
		for {
			select {
			case <-stopCh:
				return
			case w.JobsChannel <- (<-w.pool.EntryChannel):
			}
		}
	}()

	for task := range w.JobsChannel {
		if task.f == nil {
			break
		}
		task.Execute()
	}
}

func (p *Pool) Release() {
	close(stopCh)
	temp := newTask(nil)

	for i := 0; i < len(p.workers); i++ {
		p.workers[i].JobsChannel <- temp
		*(p.workers[i]) = Worker{}
	}
	p.workers = p.workers[:0]
	close(p.EntryChannel)
}

func main() {
	stopCh = make(chan struct{})

	//创建一个Task
	t := newTask(func() error {
		time.Sleep(time.Millisecond * 250)
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := newPool(3)

	//开启一个协程,不断的向Pool输送打印一条时间taks任务
	go func() {
		for {
			select {
			case <-stopCh:
				return
			case p.EntryChannel <- t:
			}
		}
	}()

	p.run()
	time.Sleep(time.Second)
	p.Release()

	time.Sleep(time.Second * 5)

	//判断goroutine是否存在泄漏, 无
	fmt.Println(runtime.NumGoroutine())
}
