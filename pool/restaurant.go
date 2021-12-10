package main

import (
	"fmt"
	"pattern-golang/concurrency/pool/workerpool"
	"sync"
	"time"
)

type SousChef struct {
	id int
}

func NewSousChef(id int) *SousChef {
	return &SousChef{id: id}
}

func (chef *SousChef) Working(taskCh chan workerpool.Task) {
	go func() {
		for task := range taskCh {
			fmt.Printf("SousChef(%d) is working task\n", chef.id)
			task.Handler(task.Data)
			fmt.Printf("SousChef(%d) has worked\n", chef.id)
		}
	}()
}

type ExecutiveChef struct {
	orderCh chan workerpool.Task
}

func NewExecutiveChef(numOfOrder int) *ExecutiveChef {
	return &ExecutiveChef{
		orderCh: make(chan workerpool.Task, numOfOrder),
	}
}

func (chef *ExecutiveChef) LaunchWorker(worker workerpool.Worker) {
	worker.Working(chef.orderCh)
}

func (chef *ExecutiveChef) ReceiveTask(task workerpool.Task) {
	select {
	case chef.orderCh <- task:
	case <-time.After(time.Second * 5):
		return
	}
}

func (chef *ExecutiveChef) Stop() {
	close(chef.orderCh)
}

func NewOrderTask(order Order, wg *sync.WaitGroup) workerpool.Task {
	return workerpool.Task{
		Data: order,
		Handler: func(i interface{}) {
			defer wg.Done()
			if o, ok := i.(Order); !ok {
				fmt.Println("order convert fail")
			} else {
				fmt.Printf("order(%d) : meal(%s)\n", o.OrderId, o.mealName)
			}
		},
	}
}

type Order struct {
	OrderId  int
	mealName string
}
