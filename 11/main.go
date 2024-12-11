package main

import (
	"fmt"
	"strconv"
)

func main() {
	inp := []int{
		4610211, 4, 0, 59, 3907, 201586, 929, 33750,
	}

	ret := Blinks(inp, 25)
	fmt.Println(len(ret))
}

func Blink(arr []int) []int {
	var res []int
	for i := range arr {
		if arr[i] == 0 {
			res = append(res, 1)
			continue
		}
		if countDigits(arr[i])%2 == 0 {
			l, r := splitNum(arr[i])
			res = append(res, l, r)
			continue
		}
		res = append(res, 2024*arr[i])
	}
	return res
}

func Blinks(arr []int, n int) []int {
	for range n {
		arr = Blink(arr)
	}
	return arr
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func splitNum(n int) (int, int) {
	numStr := strconv.Itoa(n)

	mid := len(numStr) / 2
	left := numStr[:mid]
	right := numStr[mid:]

	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)

	return leftNum, rightNum
}
