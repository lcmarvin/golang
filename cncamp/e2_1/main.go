package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go prod(ch)
	go prod(ch)
	go consume(ch)
	go consume(ch)
	time.Sleep(time.Minute)
}

func prod(ch chan<- int) {
	for {
		time.Sleep(time.Second)
		v := rand.Intn(10)
		fmt.Printf("prod%d", v)
		ch <- v
	}

}

func consume(ch <-chan int) {
	for {
		time.Sleep(time.Second)
		i := <-ch
		fmt.Printf("consume%d", i)
	}
}
