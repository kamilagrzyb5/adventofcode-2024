package main

import (
	"aoc/utils"
	"log"
	"strconv"
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

type point struct {
	i int
	j int
}

func makePoints(grid [][]int) map[point]int {
	m := map[point]int{}
	for i := range grid {
		for j := range grid[i] {
			m[point{i, j}] = grid[i][j]
		}
	}
	return m
}

func zeros(m map[point]int) map[point]int {
	result := map[point]int{}
	for k, v := range m {
		if v == 0 {
			result[k] = v
		}
	}
	return result
}

func next(points map[point]int, previous map[point]int) map[point]int {
	result := make(map[point]int)
	for k, v := range previous {
		for _, dir := range directions {
			pt := point{k.i + dir[0], k.j + dir[1]}
			if vMap, ok := points[pt]; ok && vMap == v+1 {
				result[pt] = vMap
			}
		}
	}
	return result
}

func main() {
	grid := parseInput(utils.ReadInput("10/input.txt"))

	var c int

	pts := makePoints(grid)
	start := zeros(pts)
	for k, v := range start {
		m := make(map[point]int)
		m[k] = v
		for i := 1; i <= 9; i++ {
			nxt := next(pts, m)
			m = nxt
		}
		c += len(m)
	}

	log.Printf("A is %d\n", c)
}

func parseInput(lines []string) [][]int {
	var result [][]int

	for _, line := range lines {
		var row []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result
}
