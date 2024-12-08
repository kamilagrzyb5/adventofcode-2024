package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	sum := Sum(Antinodes(MakePoints(utils.ReadInput("08/input.txt"))))
	fmt.Println(sum)
}

type Point struct {
	sign     rune
	i        int
	j        int
	antinode bool
}

func (p Point) Antena() bool {
	return p.sign != '.'
}

type Points []Point

func MakePoints(s []string) []Point {
	var points []Point
	for i, v := range s {
		for j, vv := range v {
			points = append(points, Point{
				sign: vv,
				i:    i,
				j:    j,
			})
		}
	}
	return points
}

func (points Points) Find(i, j int) (Point, int) {
	for idx, p := range points {
		if p.i == i && p.j == j {
			return p, idx
		}
	}
	return Point{}, -1
}

func Antinodes(points Points) Points {
	for _, point := range points {
		for _, p := range points {
			if point != p && point.Antena() && point.sign == p.sign {
				_, idx := points.Find(point.i+(point.i-p.i), point.j+(point.j-p.j))
				if idx != -1 {
					points[idx].antinode = true
				}
			}
		}
	}
	return points
}

func Sum(points Points) int {
	sum := 0
	for _, p := range points {
		if p.antinode {
			sum++
		}
	}
	return sum
}
