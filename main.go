package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	strs := differenceSlice(slice1, slice2)
	fmt.Println(strs)
}

func differenceSlice(s1, s2 []string) []string {
	uniqStrs := make(map[string]struct{})
	for _, str := range s2 {
		uniqStrs[str] = struct{}{}
	}
	strs := make([]string, 0, 1)
	for _, str := range s1 {
		_, ok := uniqStrs[str]
		if !ok {
			strs = append(strs, str)
		}
	}
	return strs
}
