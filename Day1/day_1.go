package day1

import (
	"aoc/utils"
	"fmt"
	"os"
)

func Day1() {
	err := utils.SetupInput("Day1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Day 1:")
	// fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func ReadInput() []string {
	file, err := os.Open("Day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var lines []string
	for {
		var line string
		_, err := fmt.Fscanln(file, &line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines
}
