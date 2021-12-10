package main

import (
	"fmt"
	"time"
)

func Order(mealNames ...string) {
	numOfMeal := len(mealNames)
	fmt.Println(numOfMeal)
	disk := make(chan Meal, numOfMeal)
	defer close(disk)
	meals := make([]Meal, numOfMeal)
	for _, name := range mealNames {
		go Cook(disk, name)
	}
	var hasErr bool
	for i := 0; i < numOfMeal; i++ {
		meal := <-disk
		if meal.Err != nil {
			fmt.Println(meal.Err)
			hasErr = true
		}
		meals[i] = meal
	}
	if !hasErr {
		for _, meal := range meals {
			fmt.Println(meal.Name)
		}
	}
}

type Meal struct {
	Name string
	Err  error
}

func Cook(disk chan<- Meal, mealName string) {
	time.Sleep(8 * time.Second)
	meal := Meal{
		mealName,
		nil,
	}
	disk <- meal
}
