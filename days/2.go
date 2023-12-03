package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run2() {
	games := convertInputToSlices()

	gameIdCount := 0

	for i, game := range games {
		if !overMaxAllowanceCheck(game) {
			gameIdCount += i + 1
		}
	}

	fmt.Println(gameIdCount)

}

func convertInputToSlices() []string {
	inputBytes, err := os.ReadFile("days/data/adventofcode.com_2023_day_2_input.txt")
	if err != nil {
		fmt.Println("could not read file")
	}

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

func overMaxAllowanceCheck(game string) bool {
	subsets := strings.Split(game, ";")
	mappedBalls := map[int]string{}

	for _, sub := range subsets {
		new := strings.Split(sub, ",")

		for _, tomap := range new {
			splitdown := strings.Split(strings.Trim(tomap, " "), " ")
			key, err := strconv.Atoi(splitdown[0])
			if err != nil {
				fmt.Println(err)
			}
			mappedBalls[key] = splitdown[1]
		}
	}

	for k, v := range mappedBalls {
		switch v {
		case "red":
			if k > 12 {
				return true
			}
		case "green":
			if k > 13 {
				return true
			}
		case "blue":
			if k > 14 {
				return true
			}
		}
	}

	return false
}
