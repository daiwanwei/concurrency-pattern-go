package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	payBack := func() (money int, err error) {
		return 100, nil
	}
	var wg sync.WaitGroup
	wg.Add(1)
	company := LoanCompany{}
	company.
		Success(func(money int) {
			fmt.Println("get money" + strconv.Itoa(money))
			wg.Done()
		}).
		Fail(func(err error) {
			fmt.Println(err)
			wg.Done()
		})
	company.Execute(payBack)
	wg.Wait()

}
