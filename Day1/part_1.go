package day1

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	numStrings = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func part1() int {
	total := 0

	for j, line := range ReadInput() {
		var (
			digit1 int
			digit2 int
		)

		fmt.Printf("line# :%d, %v, number: ", j, line)
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				digit1 = int(line[i] - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				digit2 = int(line[i] - '0')
				break
			}
		}
		total += digit1 + digit2
		fmt.Println(total)
	}
	return total
}

func part2() int {
	total := 0

	for j, line := range ReadInput() {
		var (
			minIndex = 100
			maxIndex = -1
			digit1   int
			digit2   int
		)

		for k, v := range numStrings {
			if strings.Contains(line, k) {
				if index := strings.Index(line, k); index != -1 && index < minIndex {
					digit1 = v
					minIndex = index
				}

				if index := strings.LastIndex(line, k); index != -1 && index > maxIndex {
					digit2 = v
					maxIndex = index
				}
			}
		}

		fmt.Printf("line# :%d, %v, number: ", j, line)
		for i := 0; i < minIndex; i++ {
			if unicode.IsDigit(rune(line[i])) {
				digit1 = int(line[i] - '0')
				break
			}
		}

		for i := len(line) - 1; i >= maxIndex; i-- {
			if unicode.IsDigit(rune(line[i])) {
				digit2 = int(line[i] - '0')
				break
			}
		}

		total += (digit1 * 10) + digit2
		fmt.Println(total)
	}
	return total
}
