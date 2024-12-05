package main

import (
	"aoc/utils"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFirst := utils.ReadInput("05/input1.txt")
	inputSecond := utils.ReadInput("05/input2.txt")
	orderMap := OrderMap(parseFirst(inputFirst))
	rows := parseSecond(inputSecond)
	corrects, nonCorrects := Corrects(rows, orderMap)
	fmt.Printf("count A is: %d\n", Count(corrects))
	fmt.Printf("count B is: %d\n", Count(NonCorrects(nonCorrects, orderMap)))
}

func parseFirst(s []string) [][]int {
	arr := make([][]int, len(s))
	for i := range s {
		arr[i] = make([]int, 2)
		spl := strings.Split(s[i], "|")
		arr[i][0], _ = strconv.Atoi(spl[0])
		arr[i][1], _ = strconv.Atoi(spl[1])
	}
	return arr
}

func parseSecond(s []string) [][]int {
	arr := make([][]int, len(s))
	for i := range s {
		spl := strings.Split(s[i], ",")
		arr[i] = make([]int, len(spl))
		for j := range spl {
			arr[i][j], _ = strconv.Atoi(spl[j])
		}
	}
	return arr
}

func contains(tmp []int, vm []int) bool {
	for _, t := range tmp {
		for _, v := range vm {
			if v == t {
				return false
			}
		}
	}
	return true
}

func OrderMap(arr [][]int) map[int][]int {
	res := make(map[int][]int)
	for _, v := range arr {
		res[v[0]] = append(res[v[0]], v[1])
	}
	return res
}

func NonCorrects(rows [][]int, m map[int][]int) [][]int {
	var res [][]int

	for _, row := range rows {
		for _, v := range row {
			if _, exists := m[v]; !exists {
				m[v] = make([]int, 0)
			}
		}

		arr := make([]int, len(row))
		for k, v := range m {
			var howMany int
			for _, vv := range v {
				for _, r := range row {
					if r == vv {
						howMany++
					}
				}
			}
			if slices.Contains(row, k) {
				arr[len(row)-howMany-1] = k
			}
		}
		res = append(res, arr)
	}

	return res
}

func Corrects(rows [][]int, m map[int][]int) ([][]int, [][]int) {
	var res [][]int
	var non [][]int
	for _, row := range rows {
		var tmp []int
		for _, r := range row {
			vm, ok := m[r]
			if contains(tmp, vm) || !ok {
				tmp = append(tmp, r)
			}

		}

		if reflect.DeepEqual(row, tmp) {
			res = append(res, row)
		} else {
			non = append(non, row)
		}
	}
	return res, non
}

func Count(arr [][]int) int {
	var c int
	for _, v := range arr {
		c += v[(len(v)-1)/2]
	}
	return c
}
