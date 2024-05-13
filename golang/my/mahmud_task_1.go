package my

import (
	"context"
	"fmt"
	"sync"
)

const (
	workerCount = 3
)

func workerBobr(ctx context.Context, workerId int, inputCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for bobr := range inputCh {
		fmt.Println("bobr:", bobr, "handle by worker id", workerId)
	}
}

func BobrWorkerPool() {
	data := []string{"odin bobr", "dva bobr", "tri bobr", "4otire bobr", "pyat bobr", "shest bobr"}
	inputCh := make(chan string, workerCount)
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go workerBobr(ctx, i+1, inputCh, wg)
	}

	for _, bobr := range data {
		inputCh <- bobr
	}
	close(inputCh)
	wg.Wait()
}
