package days

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type card struct {
	winning []int
	numbers []int
}

func Run4() {
	lines := convertInputToSlices(4)

	seperatedCards := []card{}

	for _, line := range lines {
		seperatedCards = append(seperatedCards, splitLines(line))
	}

	index := 0

	rounds := []int{}

	for _, card := range seperatedCards {
		roundscore := 0
		for _, ownedNum := range card.numbers {
			if slices.Contains[[]int, int](card.winning, ownedNum) {
				roundscore = roundscore * 2
				if roundscore == 0 {
					roundscore = 1
				}
			}
		}
		rounds = append(rounds, roundscore)
		index++
	}

	count := 0

	for _, r := range rounds {
		count += r
	}

}

func splitLines(line string) card {
	split := strings.Split(line, "|")
	c := card{}
	for i, split := range split {
		splitTwo := strings.Split(split, " ")
		for _, number := range splitTwo {
			number = strings.Trim(number, " ")
			if number == " " || number == "" {
				continue
			}
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println(err)

				return card{}
			}
			if i == 0 {
				c.winning = append(c.winning, num)
			} else if i == 1 {
				c.numbers = append(c.numbers, num)
			}
		}

	}

	return c
}
