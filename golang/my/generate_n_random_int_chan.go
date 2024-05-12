package my

import "math/rand"

func GenerateNRandomInt(n int) <-chan int {
	resCh := make(chan int)

	go func(out chan<- int) {
		for i := 0; i < n; i++ {
			out <- rand.Intn(100)
		}
		close(out)
	}(resCh)

	return resCh
}
