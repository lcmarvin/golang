package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	//defer close(ch)
	go consume(ch)
	prod(ch)
}

func prod(ch chan<- int) {
	for {
		time.Sleep(time.Second)
		ch <- rand.Intn(10)
	}

}

func consume(ch <-chan int) {
	for {
		time.Sleep(time.Second)
		i := <-ch
		fmt.Print(i)
	}
}
