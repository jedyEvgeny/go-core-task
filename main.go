package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()
	go func() {
		defer close(ch2)
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
	}()
	go func() {
		defer close(ch3)
		for i := 10; i < 15; i++ {
			ch3 <- i
		}
	}()

	merged := mergeChannels(ch1, ch2, ch3)

	for val := range merged {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d -> ", val)
	}
	fmt.Println("данные из канала прочитаны")
}

func mergeChannels(chls ...<-chan int) <-chan int {
	resCh := make(chan int, 5)

	go func() {
		defer close(resCh)

		for _, ch := range chls {
			for val := range ch {
				resCh <- val
			}
		}
	}()
	return resCh
}
