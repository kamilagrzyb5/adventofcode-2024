package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	b, err := os.ReadFile("03/input.txt")
	if err != nil {
		panic(err)
	}

	log.Printf("sum of muls is: %d\n", Mul(string(b)))
	log.Printf("sum of better muls is: %d\n", BetterMul(string(b)))
}

func Mul(s string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	var sum int
	for _, v := range matches {
		left, _ := strconv.Atoi(v[1])
		right, _ := strconv.Atoi(v[2])
		sum += left * right
	}

	return sum
}

func BetterMul(s string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	var sum int
	mulEnabled := true

	tokenRe := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	tokens := tokenRe.FindAllString(s, -1)

	for _, v := range tokens {
		if doRe.MatchString(v) {
			mulEnabled = true
		}

		if dontRe.MatchString(v) {
			mulEnabled = false
		}

		if mulEnabled {
			matches := re.FindStringSubmatch(v)
			if matches != nil {
				left, _ := strconv.Atoi(matches[1])
				right, _ := strconv.Atoi(matches[2])
				sum += left * right
			}
		}
	}

	return sum
}
