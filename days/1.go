package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func Run1() {
	inputs := convertInputToSlice()

	count := 0

	for _, input := range inputs {
		leftNumber, rightNumber, lowIndex, highIndex := findNumbers(input)
		leftPos, leftNum, rightPos, rightNum := findWordPositions(input)

		number1 := 0
		number2 := 0

		if leftPos < lowIndex {
			number1 = leftNum
		} else {
			number1 = leftNumber
		}

		if rightPos > highIndex {
			number2 = rightNum
		} else {
			number2 = rightNumber
		}

		finalNum, err := strconv.Atoi(fmt.Sprintf("%d%d", number1, number2))
		if err != nil {
			panic(err)
		}

		count += finalNum
	}

	fmt.Println("day 1 task 2: ", count)

}

func convertInputToSlice() []string {
	inputBytes, err := os.ReadFile("days/data/adventofcode.com_2023_day_1_input.txt")
	if err != nil {
		panic("could not read file")
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

func findNumbers(input string) (int, int, int, int) {
	sliceOfDigits := []int{}
	sliceOfPositions := []int{}

	for i := 0; i < len(input); i++ {
		index := strings.IndexFunc(input, func(r rune) bool {
			if unicode.IsDigit(r) {
				numValue, err := strconv.Atoi(fmt.Sprintf("%c", r))
				if err != nil {
					panic(err)
				}
				sliceOfDigits = append(sliceOfDigits, numValue)

				return true
			}
			return false
		})
		lastIndex := strings.LastIndexFunc(input, func(r rune) bool {
			if unicode.IsDigit(r) {
				numValue, err := strconv.Atoi(fmt.Sprintf("%c", r))
				if err != nil {
					panic(err)
				}
				sliceOfDigits = append(sliceOfDigits, numValue)

				return true
			}
			return false
		})

		sliceOfPositions = append(sliceOfPositions, index)
		sliceOfPositions = append(sliceOfPositions, lastIndex)
	}

	return sliceOfDigits[0], sliceOfDigits[len(sliceOfDigits)-1], sliceOfPositions[0], sliceOfPositions[len(sliceOfPositions)-1]
}

func findWordPositions(input string) (int, int, int, int) {
	furthestLeft := len(input)
	furthestRight := -1
	finalLeft := 0
	finalRight := 0

	for index, num := range numbers {
		indexOfEachNum := strings.Index(input, num)
		if indexOfEachNum != -1 {
			// matched on index
			if furthestLeft > indexOfEachNum {
				// if the current left position is lower than the index of a found number set the index to a new low
				furthestLeft = indexOfEachNum
				// arrays are indexed + 1 so first value is numbers[0] = 1, this converts to 1
				finalLeft = index + 1
			}
			if furthestRight < indexOfEachNum {
				furthestRight = indexOfEachNum
				finalRight = index + 1
			}
		}
	}

	return furthestLeft, finalLeft, furthestRight, finalRight
}
