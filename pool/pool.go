package main

import (
	"fmt"
	"sync"
)

type Task struct {
	Id  int
	Job func()
}

type Pool struct {
	taskQueue chan Task
	wg        sync.WaitGroup
}

func NewPool(numberWorkers int) *Pool {
	p := &Pool{
		taskQueue: make(chan Task),
	}
	p.wg.Add(numberWorkers)
	for i := 0; i < numberWorkers; i++ {
		go p.worker()
	}
	return p
}

func (p *Pool) AddTask(task Task) {
	p.taskQueue <- task
}
func (p *Pool) worker() {
	for task := range p.taskQueue {
		fmt.Printf("worker %d start task %d \n", task.Id, task.Id)
		task.Job()
		fmt.Printf("worker %d finish \n", task.Id)
	}
	p.wg.Done()
}
func (p *Pool) Wait() {
	close(p.taskQueue)
	p.wg.Wait()
}

// 实现简易pool
func main() {
	pool := NewPool(3)
	for i := 0; i < 10; i++ {
		TaskId := i
		task := Task{
			Id: TaskId,
			Job: func() {
				fmt.Printf("Task %d is running\n", TaskId)
			},
		}
		pool.AddTask(task)
	}
	pool.Wait()
}
