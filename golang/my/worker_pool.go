package my

import (
	"fmt"
	"sync"
)

func worker(id int, taskCh <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range taskCh {
		fmt.Println("worker id:", id, "res:", val*2)
		output <- val * 2
	}
}

func WorkerPool() {
	workerCount := 3
	taskCh := make(chan int, workerCount)
	resCh := make(chan int, workerCount)

	go func() {
		for num := range resCh {
			fmt.Println("read:", num)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i+1, taskCh, resCh, &wg)
	}

	ch1 := GenerateNRandomInt(7)
	ch2 := GenerateNRandomInt(8)
	ch3 := GenerateNRandomInt(5)
	mergedCh := JoinChannel(ch1, ch2, ch3)
	for res := range mergedCh {
		taskCh <- res
	}
	close(taskCh)

	wg.Wait()
	close(resCh)
}
