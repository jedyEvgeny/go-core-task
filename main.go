package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	nums, ok := intersectionSlices(a, b)
	if ok {
		fmt.Println(nums)
	}
	if !ok {
		fmt.Println("в слайсах нет одинаковых элементов")
	}
}

func intersectionSlices(a, b []int) ([]int, bool) {
	var found bool
	numCounts := make(map[int]int)
	for _, num := range a {
		numCounts[num]++
	}
	intersectionNums := make([]int, 0, 1)
	for _, el := range b {
		val, ok := numCounts[el]
		if ok && val > 0 {
			found = true
			intersectionNums = append(intersectionNums, el)
			numCounts[el]--
		}
	}
	return intersectionNums, found
}
