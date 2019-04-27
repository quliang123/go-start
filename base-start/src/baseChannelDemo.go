package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (worker *Worker) process(ints chan int) {
	for {
		select {
		case data := <-ints:
			fmt.Printf("worker %d got %d\n", worker.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		case <-time.After(time.Millisecond * 100):
			fmt.Println("timed out")
			/*这里可以留下空行以丢弃数据*/
			//default:
			//	fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}
