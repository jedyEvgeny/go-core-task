package main

import (
	"sync"
	"testing"
)

func TestMaxGorutines(t *testing.T) {
	const maxGorutines = 3

	var wg sync.WaitGroup
	var mu sync.Mutex

	testCase := []int{1, 3, 5, 10, 50}

	for _, totalTasks := range testCase {
		s := new(maxGorutines)

		var activeTasks int
		var exceededGorutins bool

		for i := 0; i < totalTasks; i++ {
			wg.Add(1)
			s.add(1)

			go func(i int) {
				defer s.done()
				defer wg.Done()

				mu.Lock()
				activeTasks++
				if activeTasks > maxGorutines {
					exceededGorutins = true
				}
				activeTasks--
				mu.Unlock()
			}(i)
		}

		wg.Wait()
		s.wait()

		if exceededGorutins {
			str := "Превышено максимальное количество горутин"
			t.Fatalf("%s\n\t%d > %d", str, activeTasks, maxGorutines)
		}
	}
}
