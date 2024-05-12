package my

import "fmt"

func PipeChan() {
	nums := make(chan int)
	sqrt := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			nums <- i
		}
		close(nums)
	}()

	go func() {
		for num := range nums {
			sqrt <- num * num
		}
		close(sqrt)
	}()

	for res := range sqrt {
		fmt.Println(res)
	}
}
