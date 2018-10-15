package main

import (
	"fmt"
	"time"
)

//定义一个任务
type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

//任务执行
func (t Task) exec() {
	t.f()
}

//并发池
type Pool struct {
	taskQue chan *Task
	poolNum int
	jobQue  chan *Task
}

func NewPool(num int) *Pool {
	p := Pool{
		taskQue: make(chan *Task),
		poolNum: num,
		jobQue:  make(chan *Task),
	}
	return &p
}

func (p *Pool) work(workID int) {
	for k := range p.jobQue {
		k.exec()
		fmt.Println("workerID :", workID, "执行完成")
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.poolNum; i++ {
		go p.work(i)
	}

	for task := range p.taskQue {
		p.jobQue <- task
	}

	close(p.jobQue)
	close(p.taskQue)
}

func main() {
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	p := NewPool(3)
	go func() {
		for {
			p.taskQue <- t
		}
	}()
	p.Run()
}
