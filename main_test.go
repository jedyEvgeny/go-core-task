package main

import "testing"

type testDifferenceSliceCase struct {
	s1, s2, expectedStrs []string
}

func TestDifferenceSlice(t *testing.T) {
	testCases := []testDifferenceSliceCase{
		{
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			[]string{},
			[]string{},
			[]string{},
		},
		{
			[]string{""},
			[]string{},
			[]string{""},
		},
		{
			[]string{},
			[]string{""},
			[]string{},
		},
		{
			[]string{"0"},
			[]string{""},
			[]string{"0"},
		},
		{
			[]string{""},
			[]string{"0"},
			[]string{""},
		},
		{
			[]string{"apple", "banana", "banana"},
			[]string{"banana"},
			[]string{"apple"},
		},
		{
			[]string{"apple", "banana"},
			[]string{"banana", "banana"},
			[]string{"apple"},
		},
		{
			[]string{"banana", "banana", "banana"},
			[]string{"banana"},
			[]string{},
		},
	}
	for _, el := range testCases {
		strs := differenceSlice(el.s1, el.s2)
		if !equalSlices(strs, el.expectedStrs) {
			t.Errorf("\nВходной слайс №1: %s\nВходной слайс №1: %s\nОжидался результат:\n\t%s\nполучили:\n\t%s\n",
				el.s1, el.s2, el.expectedStrs, strs)
		}
	}
}

func equalSlices(arr1, arrTemplate []string) bool {
	if len(arr1) != len(arrTemplate) {
		return false
	}
	//Проверка с учётом порядка элементов
	//Если порядок не важен, можно сделать проверку через карту
	for i := range arrTemplate {
		if arrTemplate[i] != arr1[i] {
			return false
		}
	}
	return true
}
