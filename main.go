package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := make(chan uint8)
	cubes := make(chan float64)

	go producer(numbers)
	go cube(numbers, cubes)
	go consumer(cubes)

	time.Sleep(5 * time.Second)
}

func producer(numbers chan<- uint8) {
	defer close(numbers)
	for {
		num := rand.Intn(100)
		fmt.Printf("Исходное число: %v\t", num)
		numbers <- uint8(num)
		time.Sleep(350 * time.Millisecond)
	}
}

func cube(numbers <-chan uint8, cubes chan<- float64) {
	defer close(cubes)
	for num := range numbers {
		numF := float64(num)
		cubed := numF * numF * numF
		fmt.Printf("Отправленное число: %v\t", cubed)
		cubes <- cubed
	}
}

func consumer(squares <-chan float64) {
	for squared := range squares {
		fmt.Printf("Получено: %v\n", squared)
	}
}
