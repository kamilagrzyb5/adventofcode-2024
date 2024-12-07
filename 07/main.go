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
	log.Printf("cnt A is: %d", Sum(Convert(inp), Evals))
	log.Printf("cnt B is: %d", Sum(Convert(inp), ExtraEvals))
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

func Sum(arr [][]int, f func([]int) []int) int {
	sum := 0
	for _, v := range arr {
		if Able(v[0], v[1:], f) {
			sum += v[0]
		}
	}
	return sum
}

func Able(result int, arr []int, f func([]int) []int) bool {
	for _, v := range f(arr) {
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

func ExtraEvals(numbers []int) []int {
	wantLen := int(math.Pow(3, float64(len(numbers)-1)))
	result := make([]int, 0, wantLen)

	result = append(result, numbers[0]+numbers[1])
	result = append(result, numbers[0]*numbers[1])
	result = append(result, concatenation(numbers[0], numbers[1]))

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
		for i := range cpRes {
			result = append(result, concatenation(cpRes[i], numbers[0]))
		}
		numbers = numbers[1:]
	}

	return result
}

func concatenation(a, b int) int {
	str := strconv.Itoa(a) + strconv.Itoa(b)

	cmb, _ := strconv.Atoi(str)

	return cmb
}
