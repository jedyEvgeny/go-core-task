package main

import "testing"

func TestCreateSlice(t *testing.T) {
	arr1 := createSlice()
	arr2 := createSlice()

	if len(arr1) != lenSlice || len(arr2) != lenSlice {
		t.Errorf("\nДлина слайса должна быть:\n\t%d\nимеется:\n\t%d\n",
			lenSlice, len(arr1))
	}

	var doubles int
	minElem, maxElem := arr1[0], arr1[0]
	for i := 0; i < lenSlice; i++ {
		if arr1[i] == arr2[i] {
			doubles++
		}

		if maxElem > arr1[i] {
			maxElem = arr1[i]
		}
		if maxElem > arr2[i] {
			maxElem = arr2[i]
		}

		if minElem < arr1[i] {
			minElem = arr1[i]
		}
		if minElem < arr2[i] {
			minElem = arr2[i]
		}
	}

	if minElem < minElemSlice {
		t.Errorf("\nМинимальный элемент слайса должен быть:\n\t%d\nимеется:\n\t%d\n",
			minElemSlice, minElem)
	}
	if maxElem > maxElemSlice {
		t.Errorf("\nМаксимальный элемент слайса должен быть:\n\t%d\nимеется:\n\t%d\n",
			maxElemSlice, maxElem)
	}

	if doubles == lenSlice {
		t.Errorf("\nОжидался случайный набор чисел\nИмеется:\n\t%d\nи\n\t%d",
			arr1, arr2)
	}
}

type evenTestCase struct {
	input, expexted []int
}

func TestSliceExample(t *testing.T) {
	testCases := []evenTestCase{
		{[]int{0}, []int{0}},
		{[]int{1, 2, 3}, []int{2}},
		{[]int{-100, 100}, []int{-100, 100}},
		{[]int{-20, -10, -90}, []int{-20, -10, -90}},
		{[]int{3, 5, 7}, []int{}},
	}
	for _, el := range testCases {
		arr := sliceExample(el.input)
		if !equalSlices(arr, el.expexted) {
			t.Errorf("\nВходной слайс: %d\nОжидался результат:\n\t%d\nполучили:\n\t%d\n",
				el.input, el.expexted, arr)
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

func TestAddElements(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5}
	num := 7
	expectedArr := []int{1, 2, 3, 4, 5, num}

	arr := addElements(testArr, num)
	if !equalSlices(arr, expectedArr) {
		t.Errorf("\nОжидалось:\n\t%d\nполучили:\n\t%d",
			expectedArr, arr)
	}
}

func TestCopySlice(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5}

	arr := copySlice(testArr)

	testArr[0] = testArr[0] + 1
	if equalSlices(arr, testArr) {
		t.Errorf("\nОжидалось:\n\t%d\nполучили:\n\t%d",
			[]int{1, 2, 3, 4, 5}, arr)
	}

	testArr = append(testArr, testArr...)
	if equalSlices(arr, testArr) {
		t.Errorf("\nОжидалось:\n\t%d\nполучили:\n\t%d",
			[]int{1, 2, 3, 4, 5}, arr)
	}
}

func TestRemoveElement(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5}
	expectedArr := []int{2, 3, 4, 5}
	idxForRemove := 0
	arrAfterDropElem := removeElement(testArr, idxForRemove)

	if !equalSlices(arrAfterDropElem, expectedArr) {
		t.Errorf("\nОжидалось:\n\t%d\nполучили:\n\t%d",
			expectedArr, arrAfterDropElem)
	}
}
