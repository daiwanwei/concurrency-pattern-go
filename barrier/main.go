package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("in")
		Order("fish", "soup")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("end")
}
