package main

import (
	"fmt"

	"math/rand"
)

const (
	lenSlice     = 10
	minElemSlice = -100
	maxElemSlice = 100
)

func main() {
	randNums := createSlice()
	fmt.Printf("Исходный слайс:\n%d\n\n", randNums)

	evenNums := sliceExample(randNums)
	fmt.Printf("Слайс чётных чисел из исходного:\n%d\n\n",
		evenNums)

	newNum := 5
	nums := addElements(evenNums, newNum)
	fmt.Printf("Слайс чётных чисел + цифра %d:\n%d\n\n",
		newNum, nums)

	copyNums := copySlice(nums)

	idxForDelete := 0
	nums = removeElement(nums, idxForDelete)

	fmt.Printf("Копия слайса чётных чисел + цифра %d:\n%d\n\n",
		newNum, copyNums)
	fmt.Printf("Слайс после удаления элемента с индексом %d:\n%d\n",
		idxForDelete, nums)
}

// createSlice генерирует слайс случайных чисел от -100 до 100.
// Логика работает для Go 1.20+, иначе будут псевдослучайные числа
func createSlice() []int {
	numbers := make([]int, lenSlice)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(maxElemSlice*2+1) + minElemSlice
	}
	return numbers
}

// sliceExample возвращает слайс чётных чисел исходного слайса
func sliceExample(nums []int) []int {
	var evenNums []int
	for _, num := range nums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}
	return evenNums
}

func addElements(nums []int, elem int) []int {
	return append(nums, elem)
}

func copySlice(nums []int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	return arr
}

func removeElement(nums []int, idx int) []int {
	nums = append(nums[:idx], nums[idx+1:]...)
	return nums
}
