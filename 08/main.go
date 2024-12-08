package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	pts := MakePoints(utils.ReadInput("08/input.txt"))
	sum := Sum(Antinodes(pts))
	sum2 := Sum(ExtraAntinodes(pts))
	fmt.Println(sum)
	fmt.Println(sum2)
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

func ExtraAntinodes(points Points) Points {
	for i, point := range points {
		for _, p := range points {
			if point.Antena() && point.sign == p.sign {
				pI, pJ := point.i+(point.i-p.i), point.j+(point.j-p.j)
				for {
					_, idx := points.Find(pI, pJ)
					if idx == -1 {
						break
					}
					if point == p {
						points[i].antinode = true
						break
					}
					points[idx].antinode = true
					pI, pJ = pI+point.i-p.i, pJ+point.j-p.j
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
