package main

import (
	"fmt"
	"time"
)

type semaphor struct {
	semaphor chan struct{}
	count    int
	doneCh   chan struct{}
}

const maxGorutines = 2

func main() {
	s := new(maxGorutines)

	for i := 0; i < 5; i++ {
		s.add(1)
		go func(i int) {
			defer s.done()
			printLoad(i)
			fmt.Printf("\rЗапуск сервиса №%d завершён!\n", i+1)
		}(i)
	}

	s.wait()
	fmt.Println("Все сервисы запущены")
}

func new(size int) *semaphor {
	return &semaphor{
		semaphor: make(chan struct{}, size),
		count:    0,
		doneCh:   make(chan struct{}),
	}
}

func (s *semaphor) add(delta int) {
	s.count += delta
	for i := 0; i < delta; i++ {
		s.semaphor <- struct{}{}
	}
}

func (s *semaphor) done() {
	<-s.semaphor
	s.count--
	if s.count == 0 {
		close(s.doneCh)
	}
}

func (s *semaphor) wait() {
	<-s.doneCh
}

func printLoad(i int) {
	symbols := []string{"/", "|", "-", "\\"}
	delay := 900
	if i%2 != 0 {
		delay = 1300
	}

	end := time.Now().Add(time.Duration(delay) * time.Millisecond)

	for time.Now().Before(end) {
		for _, el := range symbols {
			fmt.Printf("\rЗапускаем сервис: %s", el)
			time.Sleep(150 * time.Millisecond)
		}
	}
}
