package main

import "testing"

type testIntersectionSlicesCase struct {
	s1, s2, expectedStrs []int
	found                bool
}

func TestDifferenceSlice(t *testing.T) {
	testCases := []testIntersectionSlicesCase{
		{
			[]int{65, 3, 58, 678, 64},
			[]int{64, 2, 3, 43},
			[]int{64, 3},
			true,
		},
		{
			[]int{},
			[]int{},
			[]int{},
			false,
		},
		{
			[]int{0},
			[]int{},
			[]int{},
			false,
		},
		{
			[]int{},
			[]int{0},
			[]int{},
			false,
		},
		{
			[]int{-1, -2, -3, -4, -5},
			[]int{1, 2, 3, 4, 5},
			[]int{},
			false,
		},
		{
			[]int{-1, -2, -3, -4, -5},
			[]int{-1, 2, 3, 4, 5},
			[]int{-1},
			true,
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1},
			[]int{1},
			true,
		},
		{
			[]int{1},
			[]int{1, 1, 1, 1, 1},
			[]int{1},
			true,
		},
		{
			[]int{2},
			[]int{1, 1, 1, 1, 1},
			[]int{},
			false,
		},
	}
	for _, el := range testCases {
		strs, ok := intersectionSlices(el.s1, el.s2)
		if !equalSlices(strs, el.expectedStrs) || ok != el.found {
			t.Errorf("\nВходной слайс №1: %d\nВходной слайс №2: %d\nОжидался результат:\n\t%d %v\nполучили:\n\t%d %v\n",
				el.s1, el.s2, el.expectedStrs, el.found, strs, ok)
		}
	}
}

func equalSlices(arr1, arrTemplate []int) bool {
	if len(arr1) != len(arrTemplate) {
		return false
	}
	for i := range arrTemplate {
		if arrTemplate[i] != arr1[i] {
			return false
		}
	}
	return true
}
