package main

import (
	"log"
	"strconv"
)

type M map[int]int

func (m M) Add(key, value int) {
	if _, ok := m[key]; !ok {
		m[key] = 0
	}
	m[key] += value
}

func main() {
	inp := M{
		4610211: 1,
		4:       1,
		0:       1,
		59:      1,
		3907:    1,
		201586:  1,
		929:     1,
		33750:   1,
	}

	sumA := 0
	sumB := 0

	retA := Blinks(inp, 25)
	retB := Blinks(inp, 75)

	for _, v := range retA {
		sumA += v
	}

	for _, v := range retB {
		sumB += v
	}

	log.Println(sumA)
	log.Println(sumB)
}

func Blink(m M) M {
	nm := M{}

	for k, v := range m {
		if k == 0 {
			nm.Add(1, v)
			continue
		}
		if countDigits(k)%2 == 0 {
			l, r := splitNum(k)
			nm.Add(l, v)
			nm.Add(r, v)
			continue
		}
		nm.Add(k*2024, v)
	}
	return nm
}

func Blinks(m M, n int) M {
	for range n {
		m = Blink(m)
	}
	return m
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
