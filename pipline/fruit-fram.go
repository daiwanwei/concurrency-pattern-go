package main

import (
	"fmt"
	"time"
)

func SelfProducedAndSold() {
	produced := []string{"apple", "banana", "orange"}
	toSell := make(chan string, 100)
	finish := NewSellingLine(NewPackingLine(NewWashingLine(toSell)))
	go func() {
		for _, fruit := range produced {
			toSell <- fruit
		}
		close(toSell)
	}()
	for sold := range finish {
		fmt.Printf("%s sold\n", sold)
	}
}

func NewWashingLine(unWashed <-chan string) <-chan string {
	washed := make(chan string, 100)
	washing := func() {
		for fruit := range unWashed {
			fmt.Printf("%s washing\n", fruit)
			time.Sleep(time.Second)
			washed <- fruit
		}
		close(washed)
	}
	go washing()
	return washed
}

func NewPackingLine(unPacked <-chan string) <-chan string {
	packed := make(chan string, 100)
	packing := func() {
		for fruit := range unPacked {
			fmt.Printf("%s packing\n", fruit)
			time.Sleep(time.Second)
			packed <- fruit
		}
		close(packed)
	}
	go packing()
	return packed
}

func NewSellingLine(unSold <-chan string) <-chan string {
	sold := make(chan string, 100)
	selling := func() {
		for fruit := range unSold {
			fmt.Printf("%s selling\n", fruit)
			time.Sleep(time.Second)
			sold <- fruit
		}
		close(sold)
	}
	go selling()
	return sold
}
