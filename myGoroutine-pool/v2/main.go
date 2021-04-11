package main

import (
	"fmt"
	"time"
)

const (
	MaxWorker = 100
	MaxQueue  = 200
)

type Payload struct{}

type Job struct {
	PayLoad Payload
}

//全局任务队列,一个可以发送工作请求的缓冲channel
var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

//start方法开启一个worker循环，监听退出channel，可按需停止这个循环
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			// 将当前的 worker 注册到 worker 队列中
			select {
			case job := <-w.JobChannel:
				//真正业务的地方
				//模拟操作耗时
				fmt.Println("来了",job)

			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	//注册到dispatcher的worker channel池
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)		//100个缓冲
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	//开始运行n个worker
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)	//100个缓冲pool共用
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				//尝试获取一个可用的worker job channel，阻塞直到有可用的worker
				jobChan := <-d.WorkerPool
				//分发任务到worker job channel中
				jobChan <- job
			}(job)
		}
	}
}

//接收请求，把任务塞入JobQueue
func payloadHandler() {
	work := Job{PayLoad: Payload{}}
	JobQueue <- work
}

func main() {
	//通过调度器创建worker，监听来自JobQueue的任务
	d := NewDispatcher(MaxWorker)	//初始化100个缓冲的WorkerPool chan chan Job
	d.Run()
	for i:=0;i<500;i++{
		payloadHandler()
	}
	time.Sleep(5*time.Second)
}
