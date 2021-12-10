package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	netflix := NewNetflix()
	go netflix.Start()
	defer netflix.Close()
	var audiences []Audience
	for i := 0; i < 10; i++ {
		name := "ann" + strconv.Itoa(i)
		email := fmt.Sprintf("ann%d@gmail.com", i)
		a := NewAudience(i, name, email)

		audiences = append(audiences, *a)
		netflix.SubscribeChan() <- a
	}
	defer func() {
		for _, a := range audiences {
			a.Close()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 2; i++ {
			news := MovieNews{
				MovieName:   "wooo",
				Description: "dog alllllll",
			}
			netflix.PublishChan() <- news
			wg.Done()
		}
	}()
	wg.Wait()

	<-time.After(time.Second * 2)

}
