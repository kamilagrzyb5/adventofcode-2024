package main

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInput("02/input.txt")
	slice := parseInput(input)
	log.Printf("safe count is: %d\n", Safe(slice))
	log.Printf("better safe count is: %d\n", BetterSafe(slice))
}

func parseInput(input []string) [][]int {
	slice := make([][]int, len(input))
	for i, v := range input {
		nums := strings.Split(v, " ")
		row := make([]int, len(nums))
		for j, vv := range nums {
			num, _ := strconv.Atoi(vv)
			row[j] = num
		}
		slice[i] = row
	}
	return slice
}

func Safe(slice [][]int) int {
	var c int
	for _, row := range slice {
		if isSafe(row) {
			c++
		}
	}
	return c
}

func BetterSafe(slice [][]int) int {
	var c int
	for _, row := range slice {
		if isSafe(row) {
			c++
			continue
		}

		for i := range row {
			modified := append([]int{}, row[:i]...)
			modified = append(modified, row[i+1:]...)
			if isSafe(modified) {
				c++
				break
			}
		}
	}
	return c
}

func isSafe(row []int) bool {
	inc := row[1] > row[0]
	for i := range row {
		if i+1 == len(row) {
			return true
		}
		if inc {
			if row[i+1]-row[i] < 1 || row[i+1]-row[i] > 3 {
				return false
			}
		} else {
			if row[i]-row[i+1] < 1 || row[i]-row[i+1] > 3 {
				return false
			}
		}
	}
	return true
}
