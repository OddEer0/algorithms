package my

import "sync"

func BobrJoinChan(chs ...<-chan string) <-chan string {
	joinCh := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chs {
			wg.Add(1)
			go func(readCh <-chan string, wg *sync.WaitGroup) {
				defer wg.Done()
				for bobrRes := range readCh {
					joinCh <- bobrRes
				}
			}(ch, wg)
		}
		wg.Wait()
		close(joinCh)
	}()

	return joinCh
}
