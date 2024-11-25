package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	const (
		maxIdxChan = 10
		minIdxChan = 0
	)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := minIdxChan; i < maxIdxChan/2; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := maxIdxChan / 2; i < maxIdxChan; i++ {
			ch2 <- i
		}
	}()

	merged := mergeChannels(ch1, ch2)

	var count int
	valueDoubles := make(map[int]bool)

	for val := range merged {
		if val < minIdxChan || val > maxIdxChan-1 {
			t.Errorf("\nНеожиданное значение в канале\nОжидалось:\n\t[%d:%d)\nполучили:\n\t%d",
				minIdxChan, maxIdxChan-1, val)
		}

		if valueDoubles[val] {
			t.Errorf("\nОжидалось отсутствие дубликатов в канале.\nДубликат:\n\t%d",
				val)
		}

		count++
		valueDoubles[val] = true
	}

	if count != maxIdxChan {
		t.Errorf("\nНеожиданное количество значений в канале. Ожидалось:\n\t%d\nполучили:\n\t%d",
			maxIdxChan, count)
	}
}

func BenchmarkMergeChannels(b *testing.B) {
	const (
		maxIdxCh = 10_000
	)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < maxIdxCh/2; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := maxIdxCh / 2; i < maxIdxCh; i++ {
			ch2 <- i
		}
	}()

	for i := 0; i < b.N; i++ {
		merged := mergeChannels(ch1, ch2)
		for range merged {
		}
	}
}
