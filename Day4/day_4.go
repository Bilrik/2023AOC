package day4

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var (
	parsedInput []struct {
		cardNumber int
		card       [][]string
	}
)

const (
	day string = "Day4"
)

func parseInput() {
	input := utils.ReadInputLines(day)

	for i, line := range input {
		line = line[strings.Index(line, ":")+2:]
		fmt.Println(line)
		sets := strings.Split(line, " | ")
		var parsedSets [][]string
		for _, set := range sets {
			set = strings.Replace(set, "  ", " ", -1)
			if set[0] == ' ' {
				set = set[1:]
			}
			chars := strings.Split(set, " ")

			parsedSets = append(parsedSets, chars)
		}
		parsedInput = append(parsedInput, struct {
			cardNumber int
			card       [][]string
		}{i, parsedSets})
	}
	fmt.Println(parsedInput)
}

func Day4() {
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
	for i, obj := range parsedInput {
		card := obj.card
		points := 0
		for _, num := range card[0] {
			if contains(card[1], num) {
				fmt.Println("Winning Number: ", num)
				if points == 0 {
					points = 1
					continue
				}
				points *= 2
			}
		}
		fmt.Println("card: ", i, "Points: ", points)
		total += points
	}
	return total
}

func part2() int {
	cards := parsedInput
	l := len(cards)
	for i := 0; i < l; i++ {
		fmt.Println("i: ", i, "l: ", l)
		card := cards[i].card
		matches := 0
		for _, num := range card[0] {
			if contains(card[1], num) {
				matches++
			}
		}
		// fmt.Println("card# : ", cards[i].cardNumber, "Matches: ", matches)
		for matches > 0 {
			cards = append(cards, parsedInput[matches+cards[i].cardNumber])
			l++
			matches--
		}
	}
	return len(cards)
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
