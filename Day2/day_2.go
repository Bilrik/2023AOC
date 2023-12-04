package day2

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	day string = "Day2"
)

var (
	max = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func Day2() {
	err := utils.SetupInput(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 2:")
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func part1() int {
	total := 0
	for gameID, line := range utils.ReadInput(day) {
		line = line[strings.Index(line, ":")+1:]
		line = strings.TrimSpace(line)
		for _, show := range strings.Split(line, ";") {
			show = strings.TrimSpace(show)
			for _, color := range strings.Split(show, ",") {
				color = strings.TrimSpace(color)
				m := strings.Split(color, " ")
				if num, _ := strconv.Atoi(m[0]); num > max[m[1]] {
					goto NextGame
				}
			}
		}
		total += gameID + 1
	NextGame:
	}
	return total
}

func part2() int {
	power := 0
	for _, line := range utils.ReadInput(day) {
		seen := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		// trim off the game id
		line = line[strings.Index(line, ":")+1:]
		line = strings.TrimSpace(line)
		for _, show := range strings.Split(line, ";") {
			show = strings.TrimSpace(show)
			for _, color := range strings.Split(show, ",") {
				color = strings.TrimSpace(color)
				m := strings.Split(color, " ")
				if num, _ := strconv.Atoi(m[0]); num > seen[m[1]] {
					seen[m[1]] = num
				}
			}
		}
		power += seen["red"] * seen["green"] * seen["blue"]
	}
	return power
}
