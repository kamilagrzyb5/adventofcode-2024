package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.ReadFile("09/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("a: %d", checksum(move(parseStr(string(f)))))
}

func parseStr(s string) []int {
	var result []int
	var idx int
	for i, v := range s {
		vInt, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			for j := 0; j < vInt; j++ {
				result = append(result, idx)
			}
			idx++
		} else {
			for j := 0; j < vInt; j++ {
				result = append(result, -1)
			}
		}
	}
	return result
}

func move(s []int) []int {
	idx := 1
	visited := make(map[int]bool)
	for i, v := range s {
		for s[len(s)-idx] == -1 {
			visited[len(s)-idx] = true
			idx++
		}
		if v == -1 && !visited[i] {
			s[i] = s[len(s)-idx]
			s[len(s)-idx] = -1
			visited[len(s)-idx] = true
			idx++
		}
	}
	return s
}

func checksum(s []int) int {
	var sum int
	for i, v := range s {
		if v == -1 {
			break
		}
		sum += v * i
	}
	return sum
}
