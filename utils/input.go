package utils

import (
	"bufio"
	"fmt"
	"os"

	aoc "github.com/bilrik/go-aoc/pkg/api"
)

// func to setup input file
func SetupInput(folder string) error {

	if _, err := os.Stat(folder + "/input.txt"); err != nil {
		input, err := getinput()
		if err != nil {
			return err
		}
		err = os.WriteFile(folder+"/input.txt", []byte(input), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// func to download input file if needed
func getinput() (string, error) {
	client := aoc.NewClient()
	resp, err := client.GetInputData()
	if err != nil {
		return "", err
	}

	return *resp, nil
}

func ReadInputLines(folder string) []string {
	file, err := os.Open(folder + "/input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	return lines
}
