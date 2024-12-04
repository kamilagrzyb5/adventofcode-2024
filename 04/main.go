package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	arr := utils.ReadInput("04/input.txt")
	fmt.Printf("XMAS= %d\n", XMAS(arr))
	fmt.Printf("MAS= %d\n", MAS(arr))
}

func XMAS(arr []string) int {
	var c int
	for i, s := range arr {
		for j := range s {
			if j+3 < len(s) {
				if arr[i][j] == 'X' && arr[i][j+1] == 'M' && arr[i][j+2] == 'A' && arr[i][j+3] == 'S' {
					c++
				}
				if arr[i][j] == 'S' && arr[i][j+1] == 'A' && arr[i][j+2] == 'M' && arr[i][j+3] == 'X' {
					c++
				}
			}
			if i+3 < len(arr) {
				if arr[i][j] == 'X' && arr[i+1][j] == 'M' && arr[i+2][j] == 'A' && arr[i+3][j] == 'S' {
					c++
				}
				if arr[i][j] == 'S' && arr[i+1][j] == 'A' && arr[i+2][j] == 'M' && arr[i+3][j] == 'X' {
					c++
				}
			}
			if i+3 < len(arr) && j+3 < len(s) {
				if arr[i][j] == 'X' && arr[i+1][j+1] == 'M' && arr[i+2][j+2] == 'A' && arr[i+3][j+3] == 'S' {
					c++
				}
				if arr[i][j] == 'S' && arr[i+1][j+1] == 'A' && arr[i+2][j+2] == 'M' && arr[i+3][j+3] == 'X' {
					c++
				}
			}
			if i+3 < len(arr) && j-3 >= 0 {
				if arr[i][j] == 'X' && arr[i+1][j-1] == 'M' && arr[i+2][j-2] == 'A' && arr[i+3][j-3] == 'S' {
					c++
				}
				if arr[i][j] == 'S' && arr[i+1][j-1] == 'A' && arr[i+2][j-2] == 'M' && arr[i+3][j-3] == 'X' {
					c++
				}
			}
		}
	}
	return c
}

func MAS(arr []string) int {
	var c int
	for i, s := range arr {
		for j := range s {
			if i+2 < len(arr) && j+2 < len(s) {
				if arr[i][j] == 'M' && arr[i+2][j] == 'S' && arr[i+1][j+1] == 'A' && arr[i][j+2] == 'M' && arr[i+2][j+2] == 'S' {
					c++
				}
				if arr[i][j] == 'S' && arr[i+2][j] == 'M' && arr[i+1][j+1] == 'A' && arr[i][j+2] == 'S' && arr[i+2][j+2] == 'M' {
					c++
				}
				if arr[i][j] == 'S' && arr[i+2][j] == 'S' && arr[i+1][j+1] == 'A' && arr[i][j+2] == 'M' && arr[i+2][j+2] == 'M' {
					c++
				}
				if arr[i][j] == 'M' && arr[i+2][j] == 'M' && arr[i+1][j+1] == 'A' && arr[i][j+2] == 'S' && arr[i+2][j+2] == 'S' {
					c++
				}
			}
		}
	}
	return c
}
