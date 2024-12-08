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
	fmt.Printf("obstructions: %d\n", PossibleObstructions(sl))
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

func PossibleObstructions(slice [][]rune) int {
	i, j, dir := StartPos(slice)
	innerLen, outerLen := Lens(slice)
	validPositions := 0

	for x := 0; x < innerLen; x++ {
		for y := 0; y < outerLen; y++ {
			if slice[x][y] != '.' || (x == i && y == j) {
				continue
			}

			slice[x][y] = '#'

			if CausesLoop(slice, i, j, dir) {
				validPositions++
			}

			slice[x][y] = '.'
		}
	}

	return validPositions
}

func CausesLoop(slice [][]rune, startX, startY int, startDir rune) bool {
	visited := map[p]map[rune]bool{}
	i, j, dir := startX, startY, startDir
	innerLen, outerLen := Lens(slice)

	for i >= 0 && i < innerLen && j >= 0 && j < outerLen {
		pos := p{i, j}

		if _, exists := visited[pos]; exists {
			if visited[pos][dir] {
				return true
			}
		} else {
			visited[pos] = map[rune]bool{}
		}
		visited[pos][dir] = true

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

	return false
}
