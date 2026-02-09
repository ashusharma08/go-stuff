package main

import (
	"fmt"
)

func main() {
	i := 0
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	doneChan := make(chan bool)
	go func() {
		for range 50 {
			odd(&i, evenChan, oddChan, doneChan)
		}
	}()
	go func() {
		for range 50 {
			even(&i, evenChan, oddChan, doneChan)
		}

	}()
	evenChan <- true
	<-doneChan
}

func odd(i *int, evenChan chan<- bool, oddChan <-chan bool, doneChan chan bool) {
	<-oddChan
	if *i%2 != 0 {
		fmt.Println("odd ", *i)
		*i++
		val := *i

		if val > 99 {
			doneChan <- true
		}
	}
	evenChan <- true
}

func even(i *int, evenChan <-chan bool, oddChan chan<- bool, doneChan chan bool) {
	<-evenChan
	if *i%2 == 0 {
		fmt.Println("even ", *i)
		*i++
		val := *i

		if val > 99 {
			doneChan <- true
		}
	}
	oddChan <- true
}
