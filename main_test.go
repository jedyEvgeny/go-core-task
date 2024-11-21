package main

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

func TestGenerateRandomNumbers(t *testing.T) {
	randomNumbers := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go generateRandomNumbers(randomNumbers, &wg)

	count, minNum, maxNum := 0, 0, 0
	for val := range randomNumbers {
		count++
		if minNum > val {
			minNum = val
		}
		if maxNum > val {
			maxNum = val
		}
	}

	wg.Wait()

	if count != numCount {
		t.Errorf("\nОжидалось:\n\t%d чисел\nполучили\n\t%d чисел\n",
			numCount, count)
	}
	if minNum < 0 || maxNum > numRange-1 {
		t.Errorf("Нарушен допустимый диапазон. Ожидалось:\n\t[%d:%d)\nполучили:\n\t[%d:%d)",
			0, numRange, minNum, maxNum)
	}

}

func TestPrintRandomNumbers(t *testing.T) {
	ch := make(chan int)
	go func() {
		for _, num := range []int{1, 2, 3, 4, 5} {
			ch <- num
		}
		close(ch)
	}()

	var buf bytes.Buffer
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()
	r, out, _ := os.Pipe() // трубы для перенаправления
	os.Stdout = out        // стандартный вывод не в терминал, а в "out"

	go func() {
		defer func() { _ = out.Close() }()
		printRandomNumbers(ch)
	}()

	buf.ReadFrom(r)    // Читаем вывод функции из r
	os.Stdout = stdout // Вернули стандартный вывод

	output := buf.String()
	expectedOutput := "1 -> 2 -> 3 -> 4 -> 5 -> "

	if output != expectedOutput {
		t.Errorf("\nОжидалось:\n\t%v\nПолучили\n\t%v\n",
			output, expectedOutput)
	}
}
