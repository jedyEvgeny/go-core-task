package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

const errExpected = "\nОжидалось:\n\t%v\nПолучили\n\t%v\n"

func TestCreateVariables(t *testing.T) {
	testV := &variable{
		numDec:      15,
		numOct:      13,
		numHex:      21,
		numFloating: 15,
		str:         "15.0015",
		ok:          true,
		cpx:         15 + 15i,
	}
	v := createVariables()

	if testV.numDec != v.numDec {
		v1, v2 := testV.numDec, v.numDec
		t.Errorf(errExpected, v1, v2)
	}

	if testV.numOct != v.numOct {
		v1, v2 := testV.numOct, v.numOct
		t.Errorf(errExpected, v1, v2)
	}

	if testV.numHex != v.numHex {
		v1, v2 := testV.numHex, v.numHex
		t.Errorf(errExpected, v1, v2)
	}

	if testV.numFloating != v.numFloating {
		v1, v2 := testV.numFloating, v.numFloating
		t.Errorf(errExpected, v1, v2)
	}

	if testV.str != v.str {
		v1, v2 := testV.str, v.str
		t.Errorf(errExpected, v1, v2)
	}

	if testV.ok != v.ok {
		v1, v2 := testV.ok, v.ok
		t.Errorf(errExpected, v1, v2)
	}

	if testV.cpx != v.cpx {
		v1, v2 := testV.cpx, v.cpx
		t.Errorf(errExpected, v1, v2)
	}
}

func TestDefineTypes(t *testing.T) {
	v := &variable{}

	var buf bytes.Buffer
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	r, out, _ := os.Pipe() // трубы для перенаправления
	os.Stdout = out        // стандартный вывод не в терминал, а в "out"

	go func() {
		defer out.Close()
		v.defineTypes()
	}()

	buf.ReadFrom(r)    // Читаем вывод функции из r
	os.Stdout = stdout // Вернули стандартный вывод

	output := buf.String()
	expectedOutput := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n",
		"Тип переменной `numDec`: int",
		"Тип переменной `numOct`: int",
		"Тип переменной `numHex`: int",
		"Тип переменной `numFloating`: float64",
		"Тип переменной `str`: string",
		"Тип переменной `ok`: bool",
		"Тип переменной `cpx`: complex64",
	)

	if output != expectedOutput {
		t.Errorf(errExpected, output, expectedOutput)
	}
}

func TestCombineVariable(t *testing.T) {
	expectedStr := "15151515.00000015.0015true(15+15i)"
	v := &variable{
		numDec:      15,
		numOct:      13,
		numHex:      21,
		numFloating: 15,
		str:         "15.0015",
		ok:          true,
		cpx:         15 + 15i,
	}
	str := v.combineVariable()

	if expectedStr != str {
		t.Errorf(errExpected, expectedStr, str)
	}
}

func TestCreateSliceRune(t *testing.T) {
	str := "Привет, гофер!"
	runs := createSliceRune(str)

	if str != string(runs) {
		t.Errorf(errExpected, []rune(str), runs)
	}
}

func TestHashRuns(t *testing.T) {
	expectedHash := "a05865a825b7b242155f7edd06c55a901208e3bdfa22bccf36d1460772bdc884"
	input := "hello"
	actualHash, err := hashRuns([]rune(input))
	if err != nil {
		log.Fatal(err)
	}

	if actualHash != expectedHash {
		t.Errorf(errExpected, expectedHash, actualHash)
	}
}

func TestCreateStrWithSalt(t *testing.T) {
	expectedStr := "he" + salt + "llo"

	in := "hello"
	str := createStrWithSalt([]rune(in))

	if str != expectedStr {
		t.Errorf(errExpected, expectedStr, str)
	}
}
