package main

import (
	"fmt"
	"sync"
)

func main() {
	receiver := make(chan int, 20)
	// Create a dummny producer
	go func (){
		for i:= 1; i<=20; i++{
			receiver <- i
		}
	}()

	var wg sync.WaitGroup
	// Count if receive 10 msg then print
	receiverCount := make(chan int, 1)

	// Create 3 go routines to recevie msg
	for i:=1; i<= 3; i++ {
		wg.Add(1)
		go handleReceive(wg, receiver, receiverCount)
	}
	
	wg.Wait()
}

func handleReceive(wg sync.WaitGroup, receiver <-chan int, receiverCount chan int) {
	for {
		select {
		case val := <-receiver:
			c := <-receiverCount
			if (c >= 10 ) {
				fmt.Println(c)
				break
			}
			receiverCount <- 1
			fmt.Println(val)
		}
	}
	wg.Done()
}