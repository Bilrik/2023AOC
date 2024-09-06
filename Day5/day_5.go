package day4

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type seedInfo struct {
	seedNumber  int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
}

var (
	parsedInput seedInfo
	generalInfo struct {
		seedToSoil         [][]int
		soilToFertelizer   [][]int
		fertelizerToWater  [][]int
		waterToLight       [][]int
		lightToTemp        [][]int
		tempToHumidity     [][]int
		humidityToLocation [][]int
	}
)

const (
	day string = "Day5"
)

func readInput(folder string) string {
	file, err := os.Open(folder + "/input.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var contents string
	for scanner.Scan() {
		contents += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return ""
	}

	return contents
}

func parseInput() {
	input := readInput(day)

	sections := strings.Split(input, "\n\n")
	for _, section := range sections {
		parts := strings.Split(section, ":\n")
		for i, part := range parts {
			if strings.Contains(part, ":") {
				parts[i] = strings.Split(part, ": ")[1]
			}
			fmt.Println(part[i])
		}
	}
}

func Day() {
	err := utils.SetupInput(day)
	if err != nil {
		fmt.Println(err)
		return
	}
	parseInput()

	fmt.Printf("%s:\n", day)
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
