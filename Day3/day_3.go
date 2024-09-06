package day3

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	parsedInput [][]string
)

const (
	day string = "Day3"
)

func Day3() {
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
	total := 0
	for row, line := range parsedInput {
		for col, char := range line {
			if char == "." || unicode.IsDigit(rune(char[0])) || unicode.IsLetter(rune(char[0])) {
				continue
			}

			total += getTotalPartNums(row, col)
		}
	}
	return total
}

func part2() int {
	total := 0
	for row, line := range parsedInput {
		for col, char := range line {
			if char == "." || unicode.IsDigit(rune(char[0])) || unicode.IsLetter(rune(char[0])) {
				continue
			}
			if parsedInput[row][col] == "*" {
				if isGear(row, col) {
					// fmt.Println("Found gear at row: ", row, " col: ", col, " symbol: ", char)
					total += getGearRatio(row, col)
				}
			}
		}
	}
	return total
}

func parseInput() {
	input := utils.ReadInputLines(day)

	for _, line := range input {
		fmt.Println(line)
		chars := strings.Split(line, "")
		parsedInput = append(parsedInput, chars)
	}

	fmt.Println(parsedInput)
}

func getTotalPartNums(row, col int) int {
	total := sum(getTopPartNums(row, col)) + sum(getBottomPartNums(row, col)) + sum(getLeftPartNums(row, col)) + sum(getRightPartNums(row, col))
	return total
}

func sum(partNumbers []int) int {
	total := 0
	for _, num := range partNumbers {
		total += num
	}
	return total
}

func isGear(row, col int) bool {
	topPartNums := getTopPartNums(row, col)
	bottomPartNums := getBottomPartNums(row, col)
	leftPartNums := getLeftPartNums(row, col)
	rightPartNums := getRightPartNums(row, col)

	return len(topPartNums)+len(bottomPartNums)+len(leftPartNums)+len(rightPartNums) == 2
}

func getGearRatio(row, col int) int {
	ratio := 1
	topPartNums := getTopPartNums(row, col)
	bottomPartNums := getBottomPartNums(row, col)
	leftPartNums := getLeftPartNums(row, col)
	rightPartNums := getRightPartNums(row, col)

	for _, num := range topPartNums {
		ratio *= num
	}
	for _, num := range bottomPartNums {
		ratio *= num
	}
	for _, num := range leftPartNums {
		ratio *= num
	}
	for _, num := range rightPartNums {
		ratio *= num
	}

	return ratio
}

func getTopPartNums(row, col int) []int {
	partNumbers := make([]int, 0, 2)
	row--
	if row < 0 {
		return partNumbers
	}
	canidate := ""
	negCol := col - 1
	for negCol >= 0 && unicode.IsDigit(rune(parsedInput[row][negCol][0])) {
		canidate = parsedInput[row][negCol] + canidate
		negCol--
	}
	if unicode.IsDigit(rune(parsedInput[row][col][0])) {
		canidate += parsedInput[row][col]
	} else if canidate != "" {
		num, err := strconv.Atoi(canidate)
		if err != nil {
			panic(err)
		}
		partNumbers = append(partNumbers, num)
		canidate = ""
	}
	posCol := col + 1
	for posCol < len(parsedInput[row]) && unicode.IsDigit(rune(parsedInput[row][posCol][0])) {
		canidate = canidate + parsedInput[row][posCol]
		posCol++
	}
	if canidate != "" {
		num, err := strconv.Atoi(canidate)
		if err != nil {
			panic(err)
		}
		partNumbers = append(partNumbers, num)
	}
	return partNumbers
}

func getBottomPartNums(row, col int) []int {
	partNumbers := make([]int, 0, 2)
	row++
	if row >= len(parsedInput) {
		return partNumbers
	}
	canidate := ""
	negCol := col - 1
	for negCol >= 0 && unicode.IsDigit(rune(parsedInput[row][negCol][0])) {
		canidate = parsedInput[row][negCol] + canidate
		negCol--
	}
	if unicode.IsDigit(rune(parsedInput[row][col][0])) {
		canidate += parsedInput[row][col]
	} else if canidate != "" {
		num, err := strconv.Atoi(canidate)
		if err != nil {
			panic(err)
		}
		partNumbers = append(partNumbers, num)
		canidate = ""
	}
	posCol := col + 1
	for posCol < len(parsedInput[row]) && unicode.IsDigit(rune(parsedInput[row][posCol][0])) {
		canidate = canidate + parsedInput[row][posCol]
		posCol++
	}
	if canidate != "" {
		num, err := strconv.Atoi(canidate)
		if err != nil {
			panic(err)
		}
		partNumbers = append(partNumbers, num)
	}

	return partNumbers
}

func getLeftPartNums(row, col int) []int {
	partNumbers := make([]int, 0, 1)
	col--
	if col < 0 {
		return partNumbers
	}
	canidate := ""
	for col >= 0 && unicode.IsDigit(rune(parsedInput[row][col][0])) {
		canidate = parsedInput[row][col] + canidate
		col--
	}
	if canidate == "" {
		return partNumbers
	}
	num, err := strconv.Atoi(canidate)
	if err != nil {
		panic(err)
	}
	partNumbers = append(partNumbers, num)
	return partNumbers
}

func getRightPartNums(row, col int) []int {
	partNumbers := make([]int, 0, 1)
	col++
	if col >= len(parsedInput[row]) {
		return partNumbers
	}
	canidate := ""
	for col < len(parsedInput[row]) && unicode.IsDigit(rune(parsedInput[row][col][0])) {
		canidate = canidate + parsedInput[row][col]
		col++
	}
	if canidate == "" {
		return partNumbers
	}
	num, err := strconv.Atoi(canidate)
	if err != nil {
		panic(err)
	}
	partNumbers = append(partNumbers, num)
	return partNumbers
}
