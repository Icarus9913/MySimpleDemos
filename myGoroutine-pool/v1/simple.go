package main

import (
	"fmt"
	"time"
)

/*有关task任务相关定义及操作*/
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
func (t *task)Execute()  {
	t.f()	//调用任务所绑定的函数s
}

/*有关协程池的定义及操作*/
//定义池类型
type Pool struct {
	//对外接收Task的入口
	EntryChannel chan *task

	//协程池最大worker数量,限定Goroutine的个数
	worker_num int

	//协程池内部的任务就绪队列
	JobsChannel chan *task
}

//创建一个协程池
func newPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *task),
		worker_num: cap,
		JobsChannel: make(chan *task),
	}
	return &p
}

//协程池创建一个worker并且开始工作
func (p *Pool) worker(worker_ID int)  {
	//worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChannel{
		//如果拿到任务,则执行task任务
		task.Execute()
		fmt.Println("worker ID:",worker_ID," 执行完毕任务")
	}
}

//让协程池Pool开始工作
func (p *Pool)run()  {
	//1.首先根据协程池的worker数量限定,开启固定数量的Worker,
	//每一个Worker用一个Goroutine承载
	for i:=0;i<p.worker_num;i++{
		go p.worker(i)
	}

	//2.从EntryChannel协程池入口取外界传递过来的任务
	//并且将任务送进JobsChannel中
	for task := range p.EntryChannel{
		p.JobsChannel <- task
	}

	//3.执行完毕需要关闭JobsChannel
	close(p.JobsChannel)

	//4.执行完毕需要关闭EntryChannel
	close(p.EntryChannel)
}

func main()  {
	//创建一个Task
	t := newTask(func() error {
		time.Sleep(3*time.Second)
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := newPool(3)

	//开启一个协程,不断的向Pool输送打印一条时间taks任务
	go func() {
		for  {
			p.EntryChannel <- t
		}
	}()

	p.run()
}

