package day6

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

const day string = "Day6"

var (
	parsedInput [][]string
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func parseInput() {
	input := utils.ReadInputLines(day)

	for _, line := range input {
		line = strings.Split(line, ":")[1]
		line = strings.TrimSpace(line)
		for strings.Contains(line, "  ") {
			line = strings.Replace(line, "  ", " ", -1)
		}
		fmt.Println(line)

		parsedInput = append(parsedInput, strings.Split(line, " "))
	}

	for _, section := range parsedInput {
		for i, char := range section {
			if char == "" || char == " " {
				remove(section, i)
			}
		}
	}
	fmt.Println(parsedInput)
}

func Day() {
	if err := utils.SetupInput(day); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(day)
	parseInput()

	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	product := 1
	for i, length := range parsedInput[0] {
		total := 0
		raceLength, _ := strconv.Atoi(length)
		record, _ := strconv.Atoi(parsedInput[1][i])

		for j := 1; raceLength*j-j > record; j++ {
			total++
		}

		product *= total
	}

	return product
}

func part2() int {
	return 0
}
