package days

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type rows []int

func Run3() {
	lines := readInput()
	var symbolPositions []rows
	var digitPositions []rows

	for _, line := range lines {
		symbolPositions = append(symbolPositions, getSymbolPositions(line))
		digitPositions = append(digitPositions, getDigitPositions(line))
	}

	for i, symPos := range symbolPositions {
		if digitPosition[i] == symbolPosition[i] ||  digitPosition[i] == symbolPosition[i+1]  ||  digitPosition[i] == symbolPosition[i-1] {
			// this should confirm that a digit is 		
		}}
	}
}

func readInput() []string {
	inputBytes, _ := os.ReadFile("days/data/adventofcode.com_2023_day_3_input.txt")

	var lines []string

	newWordCount := 0
	for i := 0; i < len(inputBytes); i++ {
		if string(inputBytes[i]) == "\n" {
			lines = append(lines, string(inputBytes[newWordCount:i]))
			newWordCount = i
		}
	}

	return lines
}

func getSymbolPositions(line string) rows {
	f := func(r rune) bool {
		if unicode.IsSymbol(r) && r != '.' {
			return true
		}

		return false
	}

	positionsOfSymbols := rows{}

	for i := 0; i < len(line); i++ {
		pos := strings.IndexFunc(string(line[i]), f)
		if pos != -1 {
			positionsOfSymbols = append(positionsOfSymbols, i)
		}
	}

	return positionsOfSymbols
}

func getDigitPositions(line string) rows {
	f := func(r rune) bool {
		if unicode.IsDigit(r) {
			return true
		}

		return false
	}

	positionsOfDigits := rows{}

	for i := 0; i < len(line); i++ {
		pos := strings.IndexFunc(string(line[i]), f)
		if pos != -1 {
			positionsOfDigits = append(positionsOfDigits, i)
		}
	}

	return positionsOfDigits
}
