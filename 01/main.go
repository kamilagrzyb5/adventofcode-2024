package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInput("./01/input.txt")

	l, r := parseInput(lines)

	fmt.Printf("distance is: %d\n", CalculateDistance(l, r))
	fmt.Printf("similarity is: %d\n", CalculateSimilarity(l, r))
}

func parseInput(lines []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range lines {
		cols := strings.Fields(line)

		lv, _ := strconv.Atoi(cols[0])
		rv, _ := strconv.Atoi(cols[1])

		left = append(left, lv)
		right = append(right, rv)
	}

	return left, right
}

func CalculateDistance(s1, s2 []int) int {
	sort.Ints(s1)
	sort.Ints(s2)
	l := len(s1)

	var c float64

	for i := 0; i < l; i++ {
		c += math.Abs(float64(s1[i] - s2[i]))
	}

	return int(c)
}

func CalculateSimilarity(s1, s2 []int) int {
	s := 0
	for _, l := range s1 {
		c := 0
		for _, r := range s2 {
			if l == r {
				c++
			}
		}
		s += l * c
	}

	return s
}
