package main

import "testing"

type addTestCase struct {
	key         string
	val         int
	expectedMap map[string]int
}

func TestAdd(t *testing.T) {
	testCases := []addTestCase{
		{"О", 1, map[string]int{"О": 1}},
		{"Д", 2, map[string]int{"О": 1, "Д": 2}},
		{"Т", 3, map[string]int{"О": 1, "Д": 2, "Т": 3}},
		{"Ч", 4, map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4}},
		{"F", 5, map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4, "F": 5}},
	}
	si := &stringIntMap{make(map[string]int)}

	for _, el := range testCases {
		si.add(el.key, el.val)
		if !isEqualMaps(si.stringInts, el.expectedMap) {
			t.Errorf("\nОжидалась карта:\n\t%v\nимеется:\n\t%v\n",
				el.expectedMap, *si)
		}
	}
}

func isEqualMaps(m1, mExpected map[string]int) bool {
	if len(m1) != len(mExpected) {
		return false
	}
	for key, val := range mExpected {
		valM1, ok := m1[key]
		if !ok || valM1 != val {
			return false
		}
	}
	return true
}

type testRemoveCase struct {
	key         string
	initMap     map[string]int
	expextedMap map[string]int
}

func TestRevomve(t *testing.T) {
	testCases := []testRemoveCase{
		{
			"F",
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4, "F": 5},
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4},
		},
		{
			"Ч",
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4},
			map[string]int{"О": 1, "Д": 2, "Т": 3},
		},
		{
			"Т",
			map[string]int{"О": 1, "Д": 2, "Т": 3},
			map[string]int{"О": 1, "Д": 2},
		},
		{
			"Д",
			map[string]int{"О": 1, "Д": 2},
			map[string]int{"О": 1},
		},
		{
			"О",
			map[string]int{"О": 1},
			map[string]int{},
		},
	}

	for _, el := range testCases {
		si := &stringIntMap{el.initMap}

		si.remove(el.key)
		if !isEqualMaps(si.stringInts, el.expextedMap) {
			t.Errorf("\nОжидалась карта:\n\t%v\nимеется:\n\t%v\n",
				el.expextedMap, *si)
		}
	}
}

type testCopyCase struct {
	m           map[string]int
	expectedMap map[string]int
}

func TestCopy(t *testing.T) {
	testCases := []testCopyCase{
		{
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4},
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4},
		},
		{
			map[string]int{},
			map[string]int{},
		},
		{
			map[string]int{"": 0},
			map[string]int{"": 0},
		},
		{
			map[string]int{"-": -500000000000},
			map[string]int{"-": -500000000000},
		},
		{
			map[string]int{"漢字": 0},
			map[string]int{"漢字": 0},
		},
	}

	key := "testKey"
	var val int
	for _, el := range testCases {
		si := &stringIntMap{el.m}
		newM := si.copy()

		if !isEqualMaps(si.stringInts, newM) {
			t.Errorf("\nОжидалась карта:\n\t%v\nимеется:\n\t%v\n",
				el.expectedMap, newM)
		}

		si.stringInts[key] = val
		if isEqualMaps(si.stringInts, newM) {
			t.Errorf("\nКарта изменяется влед за базовой картой.\nБазовая карта:\n\t%v\nИмеется:\n\t%v\nДолжно быть:\n\t%v\n",
				si.stringInts, newM, el.expectedMap)
		}

		if len(newM) == 0 {
			continue
		}

		var randKey string
		for key := range newM {
			randKey = key
		}

		newM[randKey] = val
		if isEqualMaps(si.stringInts, newM) {
			t.Errorf("\nКарта изменяется влед за копией.\nБазовая карта:\n\t%v\nИзменённая копия:\n\t%v\nДолжно быть:\n\t%v\n",
				si.stringInts, newM, el.expectedMap)
		}
	}
}

type testIsExistsCase struct {
	key    string
	m      map[string]int
	status bool
}

func TestIsExists(t *testing.T) {
	testCases := []testIsExistsCase{
		{
			"О",
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4},
			true,
		},
		{
			"1",
			map[string]int{},
			false,
		},
		{
			"",
			map[string]int{},
			false,
		},
		{
			"",
			map[string]int{"": 0},
			true,
		},
		{
			"0",
			map[string]int{"0": 0},
			true,
		},
		{
			"0",
			map[string]int{"0": 5, "50": 0},
			true,
		},
		{
			"-",
			map[string]int{"-": -500000000000},
			true,
		},
		{
			"漢字",
			map[string]int{"漢字": 0},
			true,
		},
	}
	for _, el := range testCases {
		si := &stringIntMap{el.m}
		found := si.isExists(el.key)
		if found != el.status {
			t.Errorf("\nКлюч `%s`:\nожидалось:\n\t%v\nполучено:\n\t%v\n\t[карта: %v]",
				el.key, el.status, found, el.m)
		}
	}
}

type testGetCase struct {
	key    string
	m      map[string]int
	val    int
	status bool
}

func TestGet(t *testing.T) {
	testCases := []testGetCase{
		{
			"О",
			map[string]int{"О": 1, "Д": 2, "Т": 3, "Ч": 4},
			1,
			true,
		},
		{
			"1",
			map[string]int{},
			0,
			false,
		},
		{
			"",
			map[string]int{},
			0,
			false,
		},
		{
			"",
			map[string]int{"": 0},
			0,
			true,
		},
		{
			"0",
			map[string]int{"0": 0},
			0,
			true,
		},
		{
			"0",
			map[string]int{"0": 5, "50": 0},
			5,
			true,
		},
		{
			"-",
			map[string]int{"-": -500000000000},
			-500000000000,
			true,
		},
		{
			"漢字",
			map[string]int{"漢字": 0},
			0,
			true,
		},
	}
	for _, el := range testCases {
		si := &stringIntMap{el.m}
		val, ok := si.get(el.key)
		if ok != el.status {
			t.Errorf("\nКлюч `%s`:\nожидалось:\n\t%v\nполучено:\n\t%v\n\t[карта: %v]",
				el.key, el.status, ok, el.m)
		}
		if val != el.val {
			t.Errorf("\nКлюч `%s`:\nожидалось:\n\t%d\nполучено:\n\t%d\n\t[карта: %v]",
				el.key, val, el.val, el.m)
		}
	}
}
