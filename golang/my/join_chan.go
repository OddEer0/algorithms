package my

import "sync"

func JoinChannel(chs ...<-chan int) <-chan int {
	mergeCh := make(chan int)

	go func() {
		var wg sync.WaitGroup

		for _, ch := range chs {
			wg.Add(1)
			go func(inCh <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for res := range inCh {
					mergeCh <- res
				}
			}(ch, &wg)
		}

		wg.Wait()
		close(mergeCh)
	}()

	return mergeCh
}
