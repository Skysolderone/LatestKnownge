package main

import (
	"fmt"
	"time"
)

type Task struct {
	ExecuteTime time.Time
	Job         func()
}
type Delay struct {
	Taskquene []Task
}

func (d *Delay) addQuene(t Task) {
	d.Taskquene = append(d.Taskquene, t)
}
func (d *Delay) removeTask() {
	d.Taskquene = d.Taskquene[1:]
}
func (d *Delay) ExecuteTasks() {
	for len(d.Taskquene) > 0 {
		currentTask := d.Taskquene[0]
		if time.Now().Before(currentTask.ExecuteTime) {
			time.Sleep(currentTask.ExecuteTime.Sub(time.Now()))
		}
		currentTask.Job()
		d.removeTask()
	}

}

// 延迟队列的实现
func main() {
	fmt.Println("start")
	queue := Delay{}
	firstTask := Task{
		ExecuteTime: time.Now().Add(4 * time.Second),
		Job: func() {
			fmt.Println("EXECUTE TASK 1 AFTER ALREDY")
		},
	}
	queue.addQuene(firstTask)
	secondTask := Task{
		ExecuteTime: time.Now().Add(10 * time.Second),
		Job: func() {
			fmt.Println("EXECUTE TASK 2 AFTER DELAY")
		},
	}
	queue.addQuene(secondTask)
	queue.ExecuteTasks()
	fmt.Println("Done!")
}
