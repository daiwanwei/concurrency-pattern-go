package main

import "sync"

func main() {
	executiveChef := NewExecutiveChef(10)
	defer executiveChef.Stop()
	for i := 0; i < 3; i++ {
		sousChef := NewSousChef(i)
		executiveChef.LaunchWorker(sousChef)
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		order := Order{
			OrderId:  i,
			mealName: "haha egg",
		}
		task := NewOrderTask(order, &wg)
		executiveChef.ReceiveTask(task)
	}
	wg.Wait()
}
