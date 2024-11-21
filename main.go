package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	numCount = 10
	numRange = 100_000
)

func main() {
	fmt.Println("Генерируем случайные числа:")
	randomNumbers := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go generateRandomNumbers(randomNumbers, &wg)
	go printRandomNumbers(randomNumbers)

	wg.Wait()
	fmt.Println("конец списка")
}

func generateRandomNumbers(randomNumbers chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numCount; i++ {
		num := rand.Intn(numRange)
		randomNumbers <- num
	}
	close(randomNumbers)
}

func printRandomNumbers(randomNumbers chan int) {
	for num := range randomNumbers {
		fmt.Printf("%d -> ", num)
	}
}
