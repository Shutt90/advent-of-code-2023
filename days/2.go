package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type colour struct {
	Red   int
	Green int
	Blue  int
}

func Run2() {
	games := convertInputToSlices(2)

	gameIdCount := 0
	task2Count := 0

	for i, game := range games {
		mappedBalls := mapBalls(game, i+1)
		if !checkAllowances(mappedBalls) {
			gameIdCount += i + 1
		}
		highestMap := getHighestOfEach(mappedBalls)
		task2Count += sumUp(highestMap)
	}

	fmt.Println("day 2 task 1: ", gameIdCount)
	fmt.Println("day 2 task 2: ", task2Count)

}

func convertInputToSlices(day int) []string {
	inputBytes, _ := os.ReadFile(fmt.Sprintf("days/data/adventofcode.com_2023_day_%d_input.txt", day))

	var games []string

	newWordCount := 0
	for i := 0; i < len(inputBytes); i++ {
		if string(inputBytes[i]) == ":" {
			// remove the colon and the space
			newWordCount = i + 2
		}
		if string(inputBytes[i]) == "\n" {
			games = append(games, string(inputBytes[newWordCount:i]))
			newWordCount = i
		}
	}

	return games
}

func mapBalls(game string, index int) []colour {
	subsets := strings.Split(game, ";")
	colours := []colour{}

	for _, sub := range subsets {
		new := strings.Split(sub, ",")
		for _, tomap := range new {
			splitdown := strings.Split(strings.Trim(tomap, " "), " ")
			ballsOfColour, err := strconv.Atoi(splitdown[0])
			if err != nil {
				fmt.Println(err)
			}
			colour := colour{}
			switch splitdown[1] {
			case "red":
				colour.Red = ballsOfColour
			case "green":
				colour.Green = ballsOfColour
			case "blue":
				colour.Blue = ballsOfColour
			}
			colours = append(colours, colour)
		}
	}

	return colours
}

func checkAllowances(mappedBalls []colour) bool {
	for _, colour := range mappedBalls {
		if colour.Red > 12 {
			return true
		}
		if colour.Green > 13 {
			return true
		}
		if colour.Blue > 14 {
			return true
		}
	}

	return false
}

func getHighestOfEach(mappedBalls []colour) map[string]int {
	highestRed := 0
	highestBlue := 0
	highestGreen := 0

	for _, colour := range mappedBalls {
		if colour.Red > highestRed {
			highestRed = colour.Red
		}
		if colour.Green > highestGreen {
			highestGreen = colour.Green
		}
		if colour.Blue > highestBlue {
			highestBlue = colour.Blue
		}
	}

	newMap := map[string]int{
		"red":   highestRed,
		"green": highestGreen,
		"blue":  highestBlue,
	}

	for key, value := range newMap {
		if value == 0 {
			delete(newMap, key)
		}
	}

	return newMap
}

func sumUp(mappedBalls map[string]int) int {
	value := 1
	for _, v := range mappedBalls {
		value = v * value
	}

	return value
}
