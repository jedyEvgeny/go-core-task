package main

import "fmt"

type stringIntMap struct {
	stringInts map[string]int
}

func main() {
	si := &stringIntMap{
		make(map[string]int),
	}

	si.add("Первый", 1)
	si.add("Второй", 2)
	si.add("Третий", 3)
	fmt.Printf("Исходные данные:\n%v\n\n", *si)

	mNew := si.copy()

	si.remove("Третий")

	fmt.Printf("Копия:\n%v\n\n", mNew)
	fmt.Printf("Данные после удаления:\n%v\n\n", *si)

	key := "Второй"
	found := si.isExists(key)
	if !found {
		fmt.Printf("Ключ `%s` не найден\n", key)
	}
	if found {
		fmt.Printf("Ключ `%s` найден\n", key)
	}
	key = "Третий"
	found = si.isExists(key)
	if !found {
		fmt.Printf("Ключ `%s` не найден\n", key)
	}
	if found {
		fmt.Printf("Ключ `%s` найден\n", key)
	}

	key = "Первый"
	val, ok := si.get(key)
	if ok {
		fmt.Printf("\nЭлемент с ключём %s найден и равен: %d\n",
			key, val)
	}
	if !ok {
		fmt.Printf("\nЭлемент с кллючём %s не найден", key)
	}
}

func (si *stringIntMap) add(key string, val int) {
	si.stringInts[key] = val
}

func (si *stringIntMap) remove(key string) {
	delete(si.stringInts, key)
}

func (si *stringIntMap) copy() map[string]int {
	m := make(map[string]int)
	for key, val := range si.stringInts {
		m[key] = val
	}
	return m
}

func (si *stringIntMap) isExists(key string) bool {
	_, ok := si.stringInts[key]
	return ok
}

func (si *stringIntMap) get(key string) (int, bool) {
	val, ok := si.stringInts[key]
	return val, ok
}
