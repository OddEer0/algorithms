package my

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func reqSum(url string, resCh chan<- []byte, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	waitTime := time.Duration(rand.Intn(10))
	timer := time.NewTimer(time.Second * waitTime)
	<-timer.C
	fmt.Println(url, "Выполнен за", waitTime)
	resCh <- []byte(url)
}

func RequestsSum(urls []string) ([]byte, error) {
	resCh := make(chan []byte, 1)
	errCh := make(chan error)
	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go reqSum(url, resCh, errCh, wg)
	}

	result := []byte{}
	for i := 0; i < len(urls); i++ {
		bytes := <-resCh
		fmt.Println("Итог:", len(bytes), "байт")
		result = append(result, bytes...)
	}

	wg.Wait()

	return result, nil
}
