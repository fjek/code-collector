package main

import (
	"fmt"
	"strings"
)

// 简易版本
func MapStrUpper(strs []string, f func(s string) string) []string {
	newStrs := make([]string, len(strs))
	for i, s := range strs {
		newStrs[i] = f(s)
	}
	return newStrs
}

func MapStrsToInt(strs []string, f func(s string) int) []int {
	newArr := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		newArr[i] = f(strs[i])
	}
	return newArr
}

func reduce(strs []string, f func(s string) int) int {
	if f == nil {
		f = func(s string) int {
			return len(s)
		}
	}
	var l int = 0
	for _, s := range strs {
		l += f(s)
	}
	return l
}

func filter(arr []int, f func(i int) bool) []int {
	if f == nil {
		f = func(i int) bool {
			return i%2 == 0
		}
	}

	var newArr []int
	for _, num := range arr {
		if f(num) {
			newArr = append(newArr, num)
		}
	}
	return newArr
}

func main() {
	var list = []string{"Hao", "Chen", "MegaEase"}
	upperStrs := MapStrUpper(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("upperStrs: %v\n", upperStrs)

	ints := MapStrsToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("ints: %v\n", ints)

	length := reduce(list, nil)
	fmt.Printf("length: %v\n", length)

	filterArr := filter(ints, nil)
	fmt.Printf("filterArr: %v\n", filterArr)
}
