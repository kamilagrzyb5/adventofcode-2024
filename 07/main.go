package main

import (
	"aoc/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	inp := utils.ReadInput("07/input.txt")
	log.Printf("cnt A is: %d", Sum(Convert(inp)))
}

func Convert(s []string) [][]int {
	arr := make([][]int, len(s))
	for i, row := range s {
		parts := strings.Split(row, ":")
		result, _ := strconv.Atoi(parts[0])
		arr[i] = append(arr[i], result)
		numbers := strings.Fields(parts[1])
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			arr[i] = append(arr[i], num)
		}
	}
	return arr
}

func Sum(arr [][]int) int {
	sum := 0
	for _, v := range arr {
		if Able(v[0], v[1:]) {
			sum += v[0]
		}
	}
	return sum
}

func Able(result int, arr []int) bool {
	for _, v := range Evals(arr) {
		if v == result {
			return true
		}
	}
	return false
}

func Evals(numbers []int) []int {
	wantLen := int(math.Pow(2, float64(len(numbers)-1)))
	result := make([]int, 0, wantLen)

	result = append(result, numbers[0]+numbers[1])
	result = append(result, numbers[0]*numbers[1])

	numbers = numbers[2:]

	for len(result) < wantLen {
		cpRes := make([]int, len(result))
		copy(cpRes, result)
		for i := range result {
			result[i] += numbers[0]
		}
		for i := range cpRes {
			result = append(result, cpRes[i]*numbers[0])
		}
		numbers = numbers[1:]
	}

	return result
}
