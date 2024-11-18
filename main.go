package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

type variable struct {
	numDec      int
	numOct      int
	numHex      int
	numFloating float64
	str         string
	ok          bool
	cpx         complex64
}

const salt = "go-2024"

func main() {
	v := createVariables()
	fmt.Printf("Структура данных с переменными:\n\t%v\n\n", *v)

	v.defineTypes()

	str := v.combineVariable()
	fmt.Printf("Объединение переменных в строку:\n\t%s\n\n", str)

	runs := createSliceRune(str)
	fmt.Printf("Срез рун из строки:\n\t%v\n\n", runs)

	hRuns, err := hashRuns(runs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Хеш из рун и соли `%s`: \n\t%s\n", salt, hRuns)
}

func createVariables() *variable {
	return &variable{
		numDec:      15,
		numOct:      015,
		numHex:      0x15,
		numFloating: 15,
		str:         "15",
		ok:          true,
		cpx:         15 + 15i,
	}
}

func (v *variable) defineTypes() {
	fmt.Printf("Тип переменной `numDec`: %T\n", v.numDec)
	fmt.Printf("Тип переменной `numOct`: %T\n", v.numOct)
	fmt.Printf("Тип переменной `numHex`: %T\n", v.numHex)
	fmt.Printf("Тип переменной `numFloating`: %T\n", v.numFloating)
	fmt.Printf("Тип переменной `str`: %T\n", v.str)
	fmt.Printf("Тип переменной `ok`: %T\n", v.ok)
	fmt.Printf("Тип переменной `cpx`: %T\n", v.cpx)
	fmt.Println()
}

func (v *variable) combineVariable() string {
	return fmt.Sprintf("%d%o%x%f%s%v%v",
		v.numDec, v.numOct, v.numHex, v.numFloating, v.str, v.ok, v.cpx)
}

func createSliceRune(str string) []rune {
	return []rune(str)
}

func hashRuns(runs []rune) (string, error) {
	res := createStrWithSalt(runs)
	h := sha256.New()
	_, err := io.WriteString(h, res)
	if err != nil {
		return "", fmt.Errorf("не смогли создать хеш: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func createStrWithSalt(runs []rune) string {
	var buf bytes.Buffer
	buf.WriteString(string(runs[:len(runs)/2]))
	buf.WriteString(salt)
	buf.WriteString(string(runs[len(runs)/2:]))
	return buf.String()
}
