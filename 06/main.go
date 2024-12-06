package main

import (
	"aoc/utils"
	"fmt"
	"slices"
)

func main() {
	arr := utils.ReadInput("06/input.txt")
	sl := make([][]rune, len(arr))
	for i := range sl {
		sl[i] = []rune(arr[i])
	}

	fmt.Printf("positions: %d\n", Positions(sl))
}

var dirs = map[rune][]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

var rotateDirs = map[rune][]int{
	'^': {1, 1},
	'>': {1, -1},
	'v': {-1, -1},
	'<': {-1, 1},
}

var swapDirs = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func StartPos(slice [][]rune) (int, int, rune) {
	for i, row := range slice {
		for j, col := range row {
			if slices.Contains([]rune{'^', '>', 'v', '<'}, col) {
				return i, j, col
			}
		}
	}
	return 0, 0, ' '
}

func Lens(slice [][]rune) (int, int) {
	return len(slice), len(slice[0])
}

type p struct {
	x, y int
}

func Positions(slice [][]rune) int {
	marked := map[p]bool{}
	i, j, dir := StartPos(slice)
	innerLen, outerLen := Lens(slice)
	var c int
	for i >= 0 && i < innerLen && j >= 0 && j < outerLen {
		if slice[i][j] == '.' && !marked[p{i, j}] {
			c++
			marked[p{i, j}] = true
		}
		if slice[i][j] == '#' {
			if currDir, ok := rotateDirs[dir]; ok {
				i, j = i+currDir[0], j+currDir[1]
				dir = swapDirs[dir]
				continue
			}
		}
		if currDir, ok := dirs[dir]; ok {
			i, j = i+currDir[0], j+currDir[1]
		}
	}

	return c + 1
}
