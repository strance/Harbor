package main

import (
	"fmt"
)

type worker struct {
	c    chan int
	done chan bool
}

func main() {
	chanDemo()
}

func doWork(id int, w worker) {
	//for {
	//	n , ok := <- c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker[%d] received %c\n", id, n)
	//}
	for n := range w.c {
		fmt.Printf("worker[%d] received %c\n", id, n)
		//go func() { w.done <- true }()
		w.done <- true
	}

}

func createWorker(id int) worker {
	var w worker
	w.c = make(chan int)
	w.done = make(chan bool)
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.c <- 'a' + i
	}
	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.c <- 'A' + i
	}
	for _, worker := range workers {
		<-worker.done
	}
}
