package days

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Run1() {
	inputs := convertInputToSlice()

	count := 0
	for _, input := range inputs {
		count += findNumbers(input)
	}

	fmt.Println(count)
}

func convertInputToSlice() []string {
	inputBytes, err := os.ReadFile("days/data/adventofcode.com_2023_day_1_input.txt")
	if err != nil {
		fmt.Println("could not read file")
	}

	var inputs []string

	newWordCount := 0
	for i := 0; i < len(inputBytes); i++ {
		if string(inputBytes[i]) == "\n" {
			inputs = append(inputs, string(inputBytes[newWordCount:i]))
			newWordCount = i
		}
	}

	return inputs
}

func findNumbers(input string) int {
	start := 0
	end := len(input) - 1
	digitRegex := regexp.MustCompile(`\d`)
	runEnd := true
	runStart := true

	firstNum := 0
	lastNum := 0

	for i := 0; i < len(input); i++ {
		if runStart {
			if ok := digitRegex.Match([]byte{input[start]}); ok {
				firstNum, _ = strconv.Atoi(string(input[start]))
				runStart = false
			}
		}
		if runEnd {
			if ok := digitRegex.Match([]byte{input[end]}); ok {
				lastNum, _ = strconv.Atoi(string(input[end]))
				runEnd = false
			}
		}

		start++
		end--
	}

	number := strconv.Itoa(firstNum) + strconv.Itoa(lastNum)
	num, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}

	return num
}
